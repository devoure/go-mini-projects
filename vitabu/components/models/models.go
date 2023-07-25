package models

import (
	"vitabu/components/config"

	"github.com/jinzhu/gorm"
)

// database initialization
var myDB *gorm.DB

// Create my struct

type Books struct {
	gorm.Model          // This is a struct with predefined ID, created_at, updated_at ... fields
	BookName     string `json:"book_name"`
	AuthorName   string `json:"author_name"`
	SerialNumber string `json:"serial_number"`
}

// Create database connection, get database and migrate data

func init() {
	// create connection
	config.ConnectDB()
	// get db
	myDB = config.MyDB
	// migrate Books struct
	myDB.AutoMigrate(&Books{})
}

// CRUD functions

// Get all books from the database and return a list of books
func GetAllBooks() []Books {
	var allBooks []Books
	myDB.Find(&allBooks)
	return allBooks
}

// Create a new Book record
//  receive a struct and return a struct
func CreateBookRecord(b *Books) *Books {
	myDB.NewRecord(b) // Return true if primary key exists
	myDB.Create(&b)
	return b
}

// Get a book by ID
// receive the id of the book and return the book struct and database instance
func GetBookById(id int64) (*Books, *gorm.DB) {
	var bookRequested Books
	db := myDB.Where("ID=?", id).Find(&bookRequested)
	return &bookRequested, db
}

// Delete a book record
// receive an id and return deleted book
func DeleteBookRecord(id int64) Books {
	var book Books
	myDB.Where("ID=?", id).Delete(book)
	return book
}
