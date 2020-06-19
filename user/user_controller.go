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

func (s *Server) NewUser(ctx context.Context, in *proto.NewUserRequest) (*proto.UserResponse, error) {
	userService := NewUserService()

	user, err := userService.NewUser(in.Username, in.Email, in.FirstName, in.LastName)
	if err != nil {
		return nil, err
	}

	resp := userService.UserResponseFromModel(user)
	return resp, nil
}

func (s *Server) EditUser(ctx context.Context, in *proto.EditUserRequest) (*proto.UserResponse, error) {
	userService := NewUserService()

	user, err := userService.EditUser(in.UserID, in.User.Username, in.User.Email, in.User.FirstName, in.User.LastName)
	if err != nil {
		return nil, err
	}

	resp := userService.UserResponseFromModel(user)
	return resp, nil
}
