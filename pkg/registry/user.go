package registry

import (
	"StegoLSB/pkg/adapter/controller"
	"StegoLSB/pkg/adapter/repository"
	"StegoLSB/pkg/usecase/usecase"
)

// NewUserController conforms to interface
func (r *registry) NewUserController() controller.User {
	repo := repository.NewUserRepository(r.client)
	u := usecase.NewUserUsecase(repo)

	return controller.NewUserController(u)
}
