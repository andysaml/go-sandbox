package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	model "github.com/andysaml/go-sandbox/shop/model"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range model.Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Book{})
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book                        //инициализируем объект структуры Book
	_ = json.NewDecoder(r.Body).Decode(&book)  //парсим тело запроса и связываем её с объектом book, передаваемым по ссылке.
	book.ID = strconv.Itoa(rand.Intn(1000000)) //Дальше мы формируем случайный ID
	model.Books = append(model.Books, book)    //включаем новую книгу в массив model.Books с помощью встроенной функции append.
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range model.Books {
		if item.ID == params["id"] {
			model.Books = append(model.Books[:index], model.Books[index+1:]...)
			var book model.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			model.Books = append(model.Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(model.Books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range model.Books {
		if item.ID == params["id"] {
			model.Books = append(model.Books[:index], model.Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(model.Books)
}
