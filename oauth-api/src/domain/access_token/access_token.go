package access_token

import (
	"fmt"
	"github.com/sijanstha/oauth-api/src/utils/crypto_utils"
	"github.com/sijanstha/oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	// TODO: Validate request according to GrantType
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken(userId int64) *AccessToken {
	return &AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now())
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token!")
	}

	if at.UserId <= 0 {
		return errors.NewBadRequestError("Invalid user id!")
	}

	if at.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid client id!")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid expiration time!")
	}

	return nil
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
