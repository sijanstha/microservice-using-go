package access_token

import (
	"github.com/sijanstha/common-utils/src/utils/errors"
	"github.com/sijanstha/oauth-api/src/domain/access_token"
	"github.com/sijanstha/oauth-api/src/repository/rest"
	"strings"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(*access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type service struct {
	repository   Repository
	restUserRepo rest.RestUserRepository
}

func NewService(repo Repository, restUserRepo rest.RestUserRepository) Service {
	return &service{
		repository:   repo,
		restUserRepo: restUserRepo,
	}
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	if err := s.repository.Create(at); err != nil {
		return nil, err
	}
	return at, nil
}

func (s *service) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(token)
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token!")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
