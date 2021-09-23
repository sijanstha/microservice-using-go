package users

import (
	"github.com/sijanstha/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

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
	return nil
}
