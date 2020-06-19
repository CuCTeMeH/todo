package list

import (
	"errors"
	"github.com/jinzhu/gorm"
	"todo/model"
	"todo/proto"
	u "todo/user"
)

type ServiceI interface {
	ListingResponseFromModel(list *model.List, tasks []*model.Task) (*proto.ListResponse, error)
	GetListByID(listID string) (*model.List, error)
	GetListsForUser(userID string) ([]*model.List, error)
}

func NewListService() ServiceI {
	return &Service{}
}

type Service struct {
}

func (s Service) ListingResponseFromModel(list *model.List, tasks []*model.Task) (*proto.ListResponse, error) {
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
	user := &u.User{}
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
