package rest

import (
	"encoding/json"
	"fmt"
	"github.com/sijanstha/oauth-api/src/client/rest"
	"github.com/sijanstha/oauth-api/src/domain/users"
	"github.com/sijanstha/oauth-api/src/utils/errors"
)

var (
	userMSBaseURL = "http://localhost:8080"
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUserRepository {
	return &usersRepository{}
}

func (u usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	response, err := rest.RestClient.R().
		SetBody(users.UserLoginRequest{
			Email:    email,
			Password: password,
		}).
		Post(userMSBaseURL + "/users/login")

	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("rest client error: %s", err.Error()))
	}

	if response == nil || response.Body() == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when try to login user")
	}

	if response.StatusCode() > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal user response")
	}

	return &user, nil
}
