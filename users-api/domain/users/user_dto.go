package users

import (
	"strings"

	"github.com/sijanstha/common-utils/src/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.FirstName = strings.TrimSpace(user.FirstName)
	if user.FirstName == "" {
		return errors.NewBadRequestError("Invalid first name")
	}

	user.LastName = strings.TrimSpace(user.LastName)
	if user.LastName == "" {
		return errors.NewBadRequestError("Invalid last name")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
