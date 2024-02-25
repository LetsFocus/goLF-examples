package models

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	MathsMarks int    `json:"maths_marks"`
	Grade      string `json:"grade"`
}
