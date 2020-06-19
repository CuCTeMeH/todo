package list

import (
	"errors"
	"github.com/jinzhu/gorm"
	"todo/model"
	"todo/proto"
	"todo/task"
	u "todo/user"
)

type ListServiceI interface {
	ListingResponseFromModel(list *List, tasks []*task.Task) (*proto.ListResponse, error)
	GetListByID(listID string) (*List, error)
	GetListsForUser(userID string) ([]*List, error)
}

func NewListService() ListServiceI {
	return &ListService{}
}

type ListService struct {
}

func (s ListService) ListingResponseFromModel(list *List, tasks []*task.Task) (*proto.ListResponse, error) {
	user := &u.User{}
	err := model.Client().Model(user).Where("id = ?", list.UserID).First(&user).Error

	resp := &proto.ListResponse{
		ID:     list.UUID,
		Name:   list.Name,
		Status: list.Status,
		UserID: user.UUID,
		Tasks:  nil,
	}

	return resp, err
}

func (s ListService) GetListByID(listID string) (*List, error) {
	list := &List{}

	err := model.Client().Model(list).Where("uuid = ?", listID).First(&list).Error

	if err == gorm.ErrRecordNotFound {
		err = errors.New("record not found")
		return nil, err
	}

	return list, err
}

func (s ListService) GetListsForUser(userID string) ([]*List, error) {
	user := &u.User{}
	err := model.Client().Model(user).Where("uuid = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	uID := user.ID
	q := model.Client().Model(&List{}).Where("user_id = ?", uID)

	cnt := 0
	if err = q.Count(&cnt).Error; err != nil {
		return nil, err
	}

	lists := []*List{}

	err = q.Find(&lists).Error
	if err != nil {
		return nil, err
	}

	return lists, err
}
