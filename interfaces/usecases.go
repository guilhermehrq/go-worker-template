package interfaces

import "go-worker-template/models"

// ExampleUseCase interface of the Example usecase
type ExampleUseCase interface {
	GetExample(name string) ([]*models.Example, error)
}
