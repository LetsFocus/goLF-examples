package services

import (
	"context"
	"errors"
	"threeLayer/models"
	"threeLayer/stores"
)

type service struct {
	stores.Student
}

func New(student stores.Student) service {
	return service{student}
}

func (s service) CalculateMarks(ctx context.Context, student models.Student) (*models.Student, error) {
	if student.ID == 0 {
		return &models.Student{}, errors.New("ID is required")
	}

	if student.Name == "" {
		return &models.Student{}, errors.New("Name is required")
	}

	if student.MathsMarks == 0 {
		return &models.Student{}, errors.New("Maths Marks is required")

	}

	student.Grade = calculateGrade(student.MathsMarks)

	resp, err := s.Student.CalculateMarks(ctx, student)
	if err != nil {
		return &models.Student{}, err
	}

	return resp, nil
}

func calculateGrade(marks int) string {
	switch {
	case marks >= 90 && marks <= 100:
		return "A"
	case marks >= 80 && marks < 90:
		return "B"
	case marks >= 70 && marks < 80:
		return "C"
	case marks >= 60 && marks < 70:
		return "D"
	default:
		return "F"
	}
}
