package models

import (
	"bookstore_management_system/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
	Price  float64 `json:"price"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int) Book {
	var book Book
	db.Where("ID = ?", Id).Delete(book)
	return book
}
