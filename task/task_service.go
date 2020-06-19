package task

import (
	"todo/model"
	"todo/proto"
	u "todo/user"
)

type ServiceI interface {
	GetTaskByID() (*proto.TaskResponse, error)
	GetTasksForUser() (*proto.UserTasksResponse, error)
	GetTasksForList() (*proto.ListTasksResponse, error)
}

func NewTaskService() ServiceI {
	return &Service{}
}

type Service struct {
}

func (s Service) TaskResponseFromModel(task *Task) (*proto.TaskResponse, error) {
	user := &u.User{}
	err := model.Client().Model(user).Where("id = ?", task.UserID).First(&user).Error
	if err != nil {
		return nil, err
	}

	resp := &proto.TaskResponse{
		ID:          task.UUID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		UserID:      user.UUID,
		Deadline:    task.Deadline.Unix(),
	}

	return resp, err
}

func (s Service) GetTaskByID() (*proto.TaskResponse, error) {
	return &proto.TaskResponse{}, nil
}

func (s Service) GetTasksForUser() (*proto.UserTasksResponse, error) {
	return &proto.UserTasksResponse{}, nil
}

func (s Service) GetTasksForList() (*proto.ListTasksResponse, error) {
	return &proto.ListTasksResponse{}, nil
}
