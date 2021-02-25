package usecases_test

import (
	"errors"
	"testing"

	"go-worker-template/models"
	"go-worker-template/repositories/mocks"
	"go-worker-template/usecases"

	"github.com/stretchr/testify/assert"
)

func TestGetExample(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockExamples := make([]*models.Example, 0)

		mockExampleRepository := new(mocks.ExampleRepository)
		mockExampleRepository.On("GetExample", "test").Return(mockExamples, nil)

		exampleUseCase := usecases.NewExampleUseCase(mockExampleRepository)

		res, err := exampleUseCase.GetExample("test")

		assert.NotNil(t, res)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockExampleRepository := new(mocks.ExampleRepository)
		mockExampleRepository.On("GetExample", "test").Return(nil, errors.New(""))

		exampleUseCase := usecases.NewExampleUseCase(mockExampleRepository)

		res, err := exampleUseCase.GetExample("test")

		assert.Nil(t, res)
		assert.Error(t, err)
	})
}
