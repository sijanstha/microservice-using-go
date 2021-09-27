package services

import (
	"fmt"
	"github.com/sijanstha/domain/users"
	"github.com/sijanstha/utils/crypto"
	"github.com/sijanstha/utils/date_utils"
	"github.com/sijanstha/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	err := user.FindByEmail()
	if err == nil {
		return nil, errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
	}

	user.Password = crypto.GetMd5(user.Password)
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetTodayDateInString()
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	if &user.Id == nil || user.Id <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}

	if !isPartial {
		if err := user.Validate(); err != nil {
			return nil, err
		}
	}

	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if !isPartial {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Password = user.Password
	} else {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Password != "" {
			current.Password = user.Password
		}
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErr {

	current, err := s.GetUser(userId)
	if err != nil {
		return err
	}

	err = current.Delete()
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
