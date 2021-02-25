package interfaces

import "go-worker-template/models"

// ExampleRepository interface of the Example repository
type ExampleRepository interface {
	GetExample(name string) ([]*models.Example, error)
}
