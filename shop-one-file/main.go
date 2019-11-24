package main

// https://proglib.io/p/rest-api-go
// https://github.com/joefitzgerald/go-plus/issues/52 - как обойти автоудаление импорта (дать импорту имя, model, и исп. его перед всеми сущностями его пакета)

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return // если книга найдена, она возвращается здесь!
		}
	}
	json.NewEncoder(w).Encode(&Book{}) // возвращает пустой, не найденный экземпляр книги, просто структуру!
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book                              //инициализируем объект структуры Book
	json.NewDecoder(r.Body).Decode(&book)      //парсим тело запроса записываем по адресу только что созданной переменной book,
	book.ID = strconv.Itoa(rand.Intn(1000000)) //Дальше мы формируем случайный ID
	books = append(books, book)                //включаем новую книгу в массив books с помощью встроенной функции append.
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()
	books = append(books,
		Book{ID: "1", Title: "Война и мир",
			Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	books = append(books,
		Book{ID: "2", Title: "Преступление и наказание",
			Author: &Author{Firstname: "Федор", Lastname: "Достоевский"}})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
