package appname

import (
	"context"
	"errors"
)

var ErrResourceNotFound = errors.New("Resource does not exist/not found")

type Storage interface {
	CreateResource(context.Context, Resource) (Resource, error)
	GetResources(context.Context) ([]Resource, error)
	FindResource(context.Context, Resource) (Resource, error)
	UpdateResource(context.Context, Resource) (Resource, error)
	DeleteResource(context.Context, int) (bool, error)
}

type ResourceManager struct {
	storage Storage
}

func NewResourceManager(storage Storage) ResourceManager {
	return ResourceManager{
		storage: storage,
	}
}
