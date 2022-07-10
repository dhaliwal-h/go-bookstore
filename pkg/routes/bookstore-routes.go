package routes

import (
	"github.com/dhaliwal-h/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{bookID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
