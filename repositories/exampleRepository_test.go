package repositories_test

import (
	"errors"
	"testing"

	"go-worker-template/repositories"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetExample(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow(1, "test", "test 123")

	query := `SELECT id,
					 name,
					 description
			    FROM tb_example
			   WHERE name = $1`

	t.Run("success", func(t *testing.T) {
		mock.ExpectPrepare(query).ExpectQuery().WithArgs("test").WillReturnRows(rows)

		r := repositories.NewExampleRepository(db)

		res, err := r.GetExample("test")

		assert.NotEmpty(t, res)
		assert.NoError(t, err)
	})

	t.Run("error-scan", func(t *testing.T) {
		rowsError := sqlmock.NewRows([]string{"id", "name", "description", "error"}).AddRow(1, "test", "test 123", nil)
		mock.ExpectPrepare(query).ExpectQuery().WithArgs("test").WillReturnRows(rowsError)

		r := repositories.NewExampleRepository(db)

		res, err := r.GetExample("test")

		assert.Empty(t, res)
		assert.Error(t, err)
	})

	t.Run("error-query", func(t *testing.T) {
		mock.ExpectPrepare(query).ExpectQuery().WithArgs("test").WillReturnError(errors.New(""))

		r := repositories.NewExampleRepository(db)

		res, err := r.GetExample("test")

		assert.Empty(t, res)
		assert.Error(t, err)
	})

	t.Run("error-prepare", func(t *testing.T) {
		mock.ExpectPrepare("").ExpectQuery().WithArgs("test").WillReturnRows(rows)

		r := repositories.NewExampleRepository(db)

		res, err := r.GetExample("")

		assert.Empty(t, res)
		assert.Error(t, err)
	})
}
