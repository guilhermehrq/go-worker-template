package usecases

import (
	"go-worker-template/interfaces"
	"go-worker-template/models"
)

type exampleUseCase struct {
	exampleRepository interfaces.ExampleRepository
}

// NewExampleUseCase ...
func NewExampleUseCase(e interfaces.ExampleRepository) interfaces.ExampleUseCase {
	return &exampleUseCase{
		exampleRepository: e,
	}
}

// GetExample ...
func (u *exampleUseCase) GetExample(name string) (example []*models.Example, err error) {
	example, err = u.exampleRepository.GetExample(name)
	if err != nil {
		return nil, err
	}

	return example, nil
}
