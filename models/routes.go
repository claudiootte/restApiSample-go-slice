package models

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(http.StatusBadRequest)

	}

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(http.StatusBadRequest)

			}

			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(http.StatusBadRequest)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UTC().UnixNano())
	var book Book
	err01 := json.NewDecoder(r.Body).Decode(&book)
	if err01 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
		return
	}
	book.ID = strconv.Itoa(rand.Intn(1000000000))

	if len(book.Author) > 0 && len(book.Title) > 0 {
		books = append(books, book)
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode("Book created successfully")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(http.StatusBadRequest)

		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Author or Title cannot be empty!")
	}

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			err01 := json.NewDecoder(r.Body).Decode(&book)
			if err01 != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(http.StatusInternalServerError)
				return
			}

			book.ID = params["id"]
			books = append(books, book)
			err := json.NewEncoder(w).Encode("Book updated successfully")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(http.StatusInternalServerError)
				return
			}
			return
		}
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode("Book deleted successfully")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(http.StatusBadRequest)

	}

}
