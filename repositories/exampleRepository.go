package repositories

import (
	"database/sql"

	"go-worker-template/interfaces"
	"go-worker-template/models"
)

type exampleRepository struct {
	DBConnection *sql.DB
}

// NewExampleRepository ...
func NewExampleRepository(DBConnection *sql.DB) interfaces.ExampleRepository {
	return &exampleRepository{DBConnection}
}

func (r *exampleRepository) GetExample(name string) (examples []*models.Example, err error) {
	stmt, err := r.DBConnection.Prepare(`SELECT id,
												name,
												description
	  									   FROM tb_example
										  WHERE name = $1`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id          sql.NullInt64
			name        sql.NullString
			description sql.NullString
		)

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		example := &models.Example{}
		example.ID = id.Int64
		example.Name = name.String
		example.Description = description.String

		examples = append(examples, example)
	}

	return examples, nil
}
