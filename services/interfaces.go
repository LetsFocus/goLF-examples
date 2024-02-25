package services

import (
	"context"
	"threeLayer/models"
)

type Student interface {
	CalculateMarks(ctx context.Context, s models.Student) (*models.Student, error)
}
