package services

import (
	"fmt"
	"github.com/sijanstha/domain/users"
	"github.com/sijanstha/utils/date_utils"
	"github.com/sijanstha/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	err := user.FindByEmail()
	if err == nil {
		return nil, errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
	}

	user.DateCreated = date_utils.GetTodayDateInString()
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	if &user.Id == nil || user.Id <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if !isPartial {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
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
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {

	current, err := GetUser(userId)
	if err != nil {
		return err
	}

	err = current.Delete()
	if err != nil {
		return err
	}

	return nil
}
