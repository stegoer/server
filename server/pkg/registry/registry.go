package registry

import (
	"stegoer/ent"
	"stegoer/pkg/adapter/controller"
	"stegoer/pkg/adapter/repository"
)

type registry struct {
	client *ent.Client
}

// Registry is an interface of registry
type Registry interface {
	NewController() controller.Controller
}

// New registers entire controller with dependencies
func New(client *ent.Client) Registry {
	return &registry{
		client: client,
	}
}

// NewController generates controllers
func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		User:  repository.NewUserRepository(r.client),
		Image: repository.NewImageRepository(r.client),
	}
}
