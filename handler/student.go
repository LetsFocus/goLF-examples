
package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"threeLayer/models"
	"threeLayer/services"
)

type handler struct {
	s services.Student
}

// factory function: dependency injection
func New(student services.Student) handler {
	return handler{s: student}
}

func (h handler) PostMarks(w http.ResponseWriter, r *http.Request) {
	var s models.Student

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while reading the data"))
		return
	}
	err = json.Unmarshal(body, &s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while converting data"))
		return
	}

	_, err = h.s.CalculateMarks(r.Context(), s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
	return

}


