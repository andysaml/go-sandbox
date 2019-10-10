package main

// https://proglib.io/p/rest-api-go
// https://github.com/joefitzgerald/go-plus/issues/52 - как обойти автоудаление импорта (дать импорту имя, model, и исп. его перед всеми сущностями его пакета)

import (
	"log"
	"net/http"

	cnt "github.com/andysaml/go-sandbox/shop/controller"
	model "github.com/andysaml/go-sandbox/shop/model"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	model.Books = append(model.Books,
		model.Book{ID: "1", Title: "Война и мир",
			Author: &model.Author{Firstname: "Лев", Lastname: "Толстой"}})
	model.Books = append(model.Books,
		model.Book{ID: "2", Title: "Преступление и наказание",
			Author: &model.Author{Firstname: "Федор", Lastname: "Достоевский"}})
	r.HandleFunc("/books", cnt.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", cnt.GetBook).Methods("GET")
	r.HandleFunc("/books", cnt.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", cnt.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", cnt.DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
