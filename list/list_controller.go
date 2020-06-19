package list

import (
	"context"
	"todo/proto"
)

type Server struct {
}

func (s *Server) GetListByID(ctx context.Context, in *proto.ListRequest) (*proto.ListResponse, error) {
	listService := NewListService()

	list, err := listService.GetListByID(in.ListID)
	if err != nil {
		return nil, err
	}

	resp, err := listService.ListingResponseFromModel(list)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) GetListsForUser(ctx context.Context, in *proto.UserListsRequest) (*proto.UserListsResponse, error) {
	listService := NewListService()

	lists, err := listService.GetListsForUser(in.UserID)
	if err != nil {
		return nil, err
	}

	resp := []*proto.ListResponse{}
	for _, list := range lists {
		l, err := listService.ListingResponseFromModel(list)
		if err != nil {
			return nil, err
		}

		resp = append(resp, l)
	}

	return &proto.UserListsResponse{Lists: resp}, nil
}

func (s *Server) NewListForUser(ctx context.Context, in *proto.NewListRequest) (*proto.ListResponse, error) {
	listService := NewListService()

	list, err := listService.NewListForUser(in.UserID, in.Name, in.Status)
	if err != nil {
		return nil, err
	}

	resp, err := listService.ListingResponseFromModel(list)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
