package driver

import (
	"context"
	"errors"
)

type Store interface {
	GetDataFromKey(context context.Context, key string) string
	GetKeys(context context.Context, key string) ([]string, error)
	Increment(context context.Context, key string) error
}

// Mock Part

type StoreMock struct {
}

func (store StoreMock) Increment(context context.Context, key string) error {
	return nil
}

func (store StoreMock) GetDataFromKey(context context.Context, key string) string {
	return "4"
}

func (store StoreMock) GetKeys(context context.Context, key string) ([]string, error) {
	return []string{"example-1", "example-2"}, nil
}

type StoreErrorMock struct {
}

func (store StoreErrorMock) Increment(context context.Context, key string) error {
	return errors.New("store error")
}

func (store StoreErrorMock) GetDataFromKey(context context.Context, key string) string {
	return "4"
}

func (store StoreErrorMock) GetKeys(context context.Context, key string) ([]string, error) {
	return nil, errors.New("store error")
}

type StoreInvalidMock struct {
}

func (store StoreInvalidMock) Increment(context context.Context, key string) error {
	return nil
}

func (store StoreInvalidMock) GetDataFromKey(context context.Context, key string) string {
	return "invalid-data"
}

func (store StoreInvalidMock) GetKeys(context context.Context, key string) ([]string, error) {
	return []string{"example-1", "example-2"}, nil
}
