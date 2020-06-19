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
	listService := NewTaskService()

	tasks, err := listService.GetTasksForList(in.ListID)
	if err != nil {
		return nil, err
	}

	resp := []*proto.TaskResponse{}
	for _, task := range tasks {
		l, err := listService.TaskResponseFromModel(task)
		if err != nil {
			return nil, err
		}

		resp = append(resp, l)
	}

	return &proto.ListTasksResponse{Tasks: resp}, nil
}

func (s *Server) GetTasksForUser(ctx context.Context, in *proto.UserTasksRequest) (*proto.UserTasksResponse, error) {
	listService := NewTaskService()

	tasks, err := listService.GetTasksForUser(in.UserID)
	if err != nil {
		return nil, err
	}

	resp := []*proto.TaskResponse{}
	for _, task := range tasks {
		l, err := listService.TaskResponseFromModel(task)
		if err != nil {
			return nil, err
		}

		resp = append(resp, l)
	}

	return &proto.UserTasksResponse{Tasks: resp}, nil
}
