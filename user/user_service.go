package user

import (
	"errors"
	"github.com/jinzhu/gorm"
	"todo/model"
	"todo/proto"
)

type ServiceI interface {
	UserResponseFromModel(user *model.User) *proto.UserResponse
	GetUserByID(userID string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	NewUser(username string, email string, firstName string, lastName string) (*model.User, error)
	EditUser(userID string, username string, email string, firstName string, lastName string) (*model.User, error)
}

func NewUserService() ServiceI {
	return &Service{}
}

type Service struct {
}

func (s Service) UserResponseFromModel(user *model.User) *proto.UserResponse {
	resp := &proto.UserResponse{
		ID:        user.UUID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return resp
}

func (s Service) GetUserByID(userID string) (*model.User, error) {
	user := &model.User{}

	err := model.Client().Model(user).Where("uuid = ?", userID).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		err = errors.New("record not found")
		return nil, err
	}

	return user, err
}

func (s Service) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}

	err := model.Client().Model(user).Where("email = ?", email).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		err = errors.New("record not found")
		return nil, err
	}

	return user, err
}

func (s Service) NewUser(username string, email string, firstName string, lastName string) (*model.User, error) {
	uuid := model.UUID()

	user := &model.User{
		UUID:      uuid,
		Username:  username,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	err := model.Client().Where(user).FirstOrCreate(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s Service) EditUser(userID string, username string, email string, firstName string, lastName string) (*model.User, error) {
	user := &model.User{}
	err := model.Client().Model(user).Where("uuid = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	user.Username = username
	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName

	err = model.Client().Save(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}
