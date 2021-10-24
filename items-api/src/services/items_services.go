package services

import (
	"github.com/sijanstha/common-utils/src/utils/errors"
	"github.com/sijanstha/items-api/src/domain/items"
)


type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(string) (*items.Item, *errors.RestErr)
}

type itemsService struct{}

func NewItemService() itemServiceInterface {
	return &itemsService{}
}

func (s *itemsService) Create(item items.Item) (*items.Item, *errors.RestErr) {
	return nil, nil
}

func (s *itemsService) Get(itemId string) (*items.Item, *errors.RestErr) {
	return nil, nil
}
