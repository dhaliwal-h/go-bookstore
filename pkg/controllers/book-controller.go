package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dhaliwal-h/go-bookstore/pkg/models"
	"github.com/gorilla/mux"
)

var NewBook models.Book

var newBooks []models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newBooks = models.GetAllBooks()
	json.NewEncoder(w).Encode(newBooks)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	fmt.Printf("book id requested %v", bookId)
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	bookDetails, _ := models.GetBookById(ID)
	json.NewEncoder(w).Encode(bookDetails)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	book.CreateBook()
	fmt.Printf("%v", book)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}
	models.DeleteBook(ID)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	bookDetails, db := models.GetBookById(ID)
	if book.Author != "" {
		bookDetails.Author = book.Author
	}
	if book.Name != "" {
		bookDetails.Name = book.Name
	}
	if book.Publication != "" {
		bookDetails.Publication = book.Publication
	}
	db.Save(bookDetails)

}
