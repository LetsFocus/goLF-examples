package stores

import (
	"context"
	"database/sql"
	"threeLayer/models"
)

type store struct {
	*sql.DB
}

func New(db *sql.DB) store {
	return store{db}
}

func (s store) CalculateMarks(ctx context.Context, student models.Student) (*models.Student, error) {
	query := "INSERT INTO student (id, name, marks, grade) VALUES ($1, $2, $3, $4)"

	resp, err := s.ExecContext(ctx, query, student.ID, student.Name, student.MathsMarks, student.Grade)
	if err != nil {
		return nil, err
	}

	_, err = resp.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &student, nil
}
