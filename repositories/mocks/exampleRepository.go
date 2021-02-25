package mocks

import (
	"go-worker-template/models"

	"github.com/stretchr/testify/mock"
)

// ExampleRepository ...
type ExampleRepository struct {
	mock.Mock
}

// GetExample ...
func (r *ExampleRepository) GetExample(name string) ([]*models.Example, error) {
	args := r.Called(name)

	var r0 []*models.Example
	if rf, ok := args.Get(0).(func(string) []*models.Example); ok {
		r0 = rf(name)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]*models.Example)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
