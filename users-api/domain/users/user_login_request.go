package users

import (
	"github.com/sijanstha/utils/errors"
	"strings"
)

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserLoginRequest) Validate() *errors.RestErr {
	r.Email = strings.TrimSpace(strings.ToLower(r.Email))
	if r.Email == "" {
		return errors.NewBadRequestError("Invalid email")
	}

	if r.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}

	return nil
}
