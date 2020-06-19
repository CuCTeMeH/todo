package user

import (
	"context"
	"todo/proto"
)

type Server struct {
}

func (s *Server) GetUserByID(ctx context.Context, in *proto.UserRequest) (*proto.UserResponse, error) {
	userService := NewUserService()

	user, err := userService.GetUserByID(in.UserID)
	if err != nil {
		return nil, err
	}

	resp := userService.UserResponseFromModel(user)
	return resp, nil
}

func (s *Server) GetUserByEmail(ctx context.Context, in *proto.UserByEmailRequest) (*proto.UserResponse, error) {
	userService := NewUserService()

	user, err := userService.GetUserByEmail(in.Email)
	if err != nil {
		return nil, err
	}

	resp := userService.UserResponseFromModel(user)
	return resp, nil
}
