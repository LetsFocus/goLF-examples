package main

import (
	"github.com/LetsFocus/goLF/goLF"
	"net/http"
	"threeLayer/handler"
	"threeLayer/services"
	"threeLayer/stores"
)

//func main() {
//	store := stores.New()
//	service := services.New(store)
//	handler := handlers.New(service)
//
//	http.HandleFunc("/student", handler.PostMarks)
//
//	http.ListenAndServe(":8000", nil)
//}

func main() {
	g := goLF.New()

	store := stores.New(g.Database.Postgres)
	service := services.New(store)
	handler := handlers.New(service)

	http.HandleFunc("/student", handler.PostMarks)

	http.ListenAndServe(":8001", nil)
}
