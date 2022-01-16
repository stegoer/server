package registry

import (
	"StegoLSB/pkg/adapter/controller"
	"StegoLSB/pkg/adapter/repository"
	"StegoLSB/pkg/usecase/usecase"
)

func (r *registry) NewImageController() controller.Image {
	repo := repository.NewImageRepository(r.client)
	u := usecase.NewImageUsecase(repo)

	return controller.NewImageController(u)
}
