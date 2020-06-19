package task

import (
	"context"
	"log"
	"todo/proto"
)

type Server struct {
}

func (s *Server) getTask(ctx context.Context, in *proto.TaskRequest) (*proto.TaskResponse, error) {
	log.Printf("Get Todo list with uuid: %s", in.GetUuid())
	return &proto.TaskResponse{Name: "", Status: ""}, nil
}

func (s *Server) GetTasksForList(ctx context.Context, in *proto.ListTasksRequest) (*proto.ListTasksResponse, error) {
	log.Printf("User Id for fetching todo lists: %s", in.GetListID())
	return &proto.ListTasksResponse{}, nil
}

func (s *Server) GetTasksForUser(ctx context.Context, in *proto.UserTasksRequest) (*proto.UserTasksResponse, error) {
	log.Printf("User Id for fetching todo lists: %s", in.GetUserID())
	return &proto.UserTasksResponse{}, nil
}
