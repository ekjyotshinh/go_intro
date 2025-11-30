package controllers

import (
	"bookstore_management_system/pkg/models"
	"bookstore_management_system/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
func GetBooks(w http.ResponseWriter, r *http.Request){
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	book, db := models.GetBookById(id)
	if db.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseJSONBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var book = &models.Book{}
	utils.ParseJSONBody(r, book)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	bookDetails, db := models.GetBookById(id)
	if db.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if book.Title != "" {
		bookDetails.Title = book.Title
	}
	if book.Author != "" {
		bookDetails.Author = book.Author
	}
	if book.Publisher != "" {
		bookDetails.Publisher = book.Publisher
	}
	if book.Price != 0 {
		bookDetails.Price = book.Price
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	deletedBook := models.DeleteBook(id)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}