package list

import (
	"errors"
	"github.com/jinzhu/gorm"
	"todo/model"
	"todo/proto"
	taskService "todo/task"
)

type ServiceI interface {
	ListingResponseFromModel(list *model.List) (*proto.ListResponse, error)
	GetListByID(listID string) (*model.List, error)
	GetListsForUser(userID string) ([]*model.List, error)
	NewListForUser(userID string, name string, status string) (*model.List, error)
}

func NewListService() ServiceI {
	return &Service{}
}

type Service struct {
}

func (s Service) ListingResponseFromModel(list *model.List) (*proto.ListResponse, error) {
	user := &model.User{}
	err := model.Client().Model(user).Where("id = ?", list.UserID).First(&user).Error

	tasks := []*model.Task{}
	err = model.Client().Model(tasks).Where("list_id = ?", list.ID).Find(&tasks).Error

	taskResp := []*proto.TaskResponse{}
	for _, task := range tasks {
		t, err := taskService.NewTaskService().TaskResponseFromModel(task)
		if err != nil {
			return nil, err
		}

		taskResp = append(taskResp, t)
	}

	resp := &proto.ListResponse{
		ID:     list.UUID,
		Name:   list.Name,
		Status: list.Status,
		UserID: user.UUID,
		Tasks:  taskResp,
	}

	return resp, err
}

func (s Service) GetListByID(listID string) (*model.List, error) {
	list := &model.List{}

	err := model.Client().Model(list).Where("uuid = ?", listID).First(&list).Error

	if err == gorm.ErrRecordNotFound {
		err = errors.New("record not found")
		return nil, err
	}

	return list, err
}

func (s Service) GetListsForUser(userID string) ([]*model.List, error) {
	user := &model.User{}
	err := model.Client().Model(user).Where("uuid = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	uID := user.ID
	q := model.Client().Model(&model.List{}).Where("user_id = ?", uID)

	cnt := 0
	if err = q.Count(&cnt).Error; err != nil {
		return nil, err
	}

	lists := []*model.List{}

	err = q.Find(&lists).Error
	if err != nil {
		return nil, err
	}

	return lists, err
}

func (s Service) NewListForUser(userID string, name string, status string) (*model.List, error) {
	uuid := model.UUID()

	user := &model.User{}
	err := model.Client().Model(user).Where("uuid = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	list := &model.List{
		UUID:   uuid,
		UserID: user.ID,
		Name:   name,
		Status: status,
	}

	err = model.Client().Where(list).FirstOrCreate(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
