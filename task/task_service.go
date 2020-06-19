package task

import (
	"errors"
	"github.com/jinzhu/gorm"
	"todo/model"
	"todo/proto"
	u "todo/user"
)

type ServiceI interface {
	TaskResponseFromModel(task *model.Task) (*proto.TaskResponse, error)
	GetTaskByID(taskID string) (*model.Task, error)
	GetTasksForUser(userID string) ([]*model.Task, error)
	GetTasksForList(listID string) ([]*model.Task, error)
}

func NewTaskService() ServiceI {
	return &Service{}
}

type Service struct {
}

func (s Service) TaskResponseFromModel(task *model.Task) (*proto.TaskResponse, error) {
	user := &u.User{}
	err := model.Client().Model(user).Where("id = ?", task.UserID).First(&user).Error
	if err != nil {
		return nil, err
	}

	l := &model.List{}
	err = model.Client().Model(l).Where("id = ?", task.ListID).First(&l).Error
	if err != nil {
		return nil, err
	}

	resp := &proto.TaskResponse{
		ID:          task.UUID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		UserID:      user.UUID,
		ListID:      l.UUID,
		Deadline:    task.Deadline.Unix(),
	}

	return resp, err
}

func (s Service) GetTaskByID(taskID string) (*model.Task, error) {
	task := &model.Task{}

	err := model.Client().Model(task).Where("uuid = ?", taskID).First(&task).Error
	if err == gorm.ErrRecordNotFound {
		err = errors.New("record not found")
		return nil, err
	}

	return task, err
}

func (s Service) GetTasksForUser(userID string) ([]*model.Task, error) {
	user := &u.User{}
	err := model.Client().Model(user).Where("uuid = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	uID := user.ID
	q := model.Client().Model(&model.Task{}).Where("user_id = ?", uID)

	cnt := 0
	if err = q.Count(&cnt).Error; err != nil {
		return nil, err
	}

	tasks := []*model.Task{}

	err = q.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, err
}

func (s Service) GetTasksForList(listID string) ([]*model.Task, error) {
	l := &model.List{}
	err := model.Client().Model(l).Where("uuid = ?", listID).First(&l).Error
	if err != nil {
		return nil, err
	}

	uID := l.ID
	q := model.Client().Model(&model.Task{}).Where("list_id = ?", uID)

	cnt := 0
	if err = q.Count(&cnt).Error; err != nil {
		return nil, err
	}

	tasks := []*model.Task{}

	err = q.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, err
}
