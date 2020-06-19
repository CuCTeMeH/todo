package task

import (
	"context"
	"todo/proto"
)

type Server struct {
}

func (s *Server) GetTaskByID(ctx context.Context, in *proto.TaskRequest) (*proto.TaskResponse, error) {
	taskService := NewTaskService()

	task, err := taskService.GetTaskByID(in.TaskID)
	if err != nil {
		return nil, err
	}

	resp, err := taskService.TaskResponseFromModel(task)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) GetTasksForList(ctx context.Context, in *proto.ListTasksRequest) (*proto.ListTasksResponse, error) {
	taskService := NewTaskService()

	tasks, err := taskService.GetTasksForList(in.ListID)
	if err != nil {
		return nil, err
	}

	resp := []*proto.TaskResponse{}
	for _, task := range tasks {
		l, err := taskService.TaskResponseFromModel(task)
		if err != nil {
			return nil, err
		}

		resp = append(resp, l)
	}

	return &proto.ListTasksResponse{Tasks: resp}, nil
}

func (s *Server) GetTasksForUser(ctx context.Context, in *proto.UserTasksRequest) (*proto.UserTasksResponse, error) {
	taskService := NewTaskService()

	tasks, err := taskService.GetTasksForUser(in.UserID)
	if err != nil {
		return nil, err
	}

	resp := []*proto.TaskResponse{}
	for _, task := range tasks {
		l, err := taskService.TaskResponseFromModel(task)
		if err != nil {
			return nil, err
		}

		resp = append(resp, l)
	}

	return &proto.UserTasksResponse{Tasks: resp}, nil
}

func (s *Server) NewTask(ctx context.Context, in *proto.NewTaskRequest) (*proto.TaskResponse, error) {
	taskService := NewTaskService()

	task, err := taskService.NewTask(in.ListID, in.UserID, in.Name, in.Description, in.Status, in.Deadline)
	if err != nil {
		return nil, err
	}

	resp, err := taskService.TaskResponseFromModel(task)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) EditTask(ctx context.Context, in *proto.EditTaskRequest) (*proto.TaskResponse, error) {
	taskService := NewTaskService()

	task, err := taskService.EditTask(in.TaskID, in.Task.ListID, in.Task.UserID, in.Task.Name, in.Task.Description, in.Task.Status, in.Task.Deadline)
	if err != nil {
		return nil, err
	}

	resp, err := taskService.TaskResponseFromModel(task)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
