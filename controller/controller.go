package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/claudiootte/restApiSample-go-slice/models"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	MyRouter := mux.NewRouter()

	MyRouter.HandleFunc("/books", models.GetAllBooks).Methods("GET")
	MyRouter.HandleFunc("/books/{id}", models.GetBook).Methods("GET")
	MyRouter.HandleFunc("/books", models.CreateBook).Methods("POST")
	MyRouter.HandleFunc("/books/{id}", models.UpdateBook).Methods("PUT")
	MyRouter.HandleFunc("/books/{id}", models.DeleteBook).Methods("DELETE")

	fmt.Println("Connecting to server...")
	log.Fatal(http.ListenAndServe(":8080", MyRouter))
}
