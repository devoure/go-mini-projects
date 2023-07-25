package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"vitabu/components/models"
	"vitabu/components/utils"

	"github.com/gorilla/mux"
)

// getAllBooks controller (handler)
func GetAllBooksController(w http.ResponseWriter, r *http.Request) {
	// get all books
	allBooks := models.GetAllBooks()
	// change struct to json
	response, _ := json.Marshal(allBooks)
	// write the response
	// define the format of response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// getBookById controller (handler)
func GetBookByIdController(w http.ResponseWriter, r *http.Request) {
	// get requested id from url
	vars := mux.Vars(r)
	requestedId := vars["book_id"]
	// parse id from string to int
	bookId, err := strconv.ParseInt(requestedId, 0, 0)
	// check if parse was successful
	if err != nil {
		fmt.Println(">>> ERROR while parsing book ID")
	}
	// get book by id from database
	requestedBook, _ := models.GetBookById(bookId)
	// marshall book struct to json
	response, _ := json.Marshal(requestedBook)
	// write response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// createBookRecord controller (handlers)
func CreateBookRecordController(w http.ResponseWriter, r *http.Request) {
	// create var for new record
	newBookRecord := &models.Books{}
	// parse body contents
	utils.ParseBody(r, newBookRecord)
	// create new record in database
	books := models.CreateBookRecord(newBookRecord)
	// write response
	response, _ := json.Marshal(books)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// delete book record controller
func DeleteBookRecordController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedId := vars["book_id"]
	// parse id
	bookId, err := strconv.ParseInt(requestedId, 0, 0)
	if err != nil {
		fmt.Println(">>> ERROR while parsing book ID")
	}
	// delete book with specified id from database
	books := models.DeleteBookRecord(bookId)
	// marshall books to json
	response, _ := json.Marshal(books)
	// write the response
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// update book record controller (handlers)
func UpdateBookRecordController(w http.ResponseWriter, r *http.Request) {
	newBookRecord := &models.Books{}
	// parse body contents from request
	utils.ParseBody(r, newBookRecord)
	// get requested id from url
	vars := mux.Vars(r)
	requestedId := vars["book_id"]
	// parse id
	bookId, err := strconv.ParseInt(requestedId, 0, 0)
	if err != nil {
		fmt.Println("ERROR while parsing book id")
	}
	// get book by id
	requestedBook, db := models.GetBookById(bookId)
	// update records
	// update only fields that need to be updated
	if newBookRecord.BookName != "" {
		requestedBook.BookName = newBookRecord.BookName
	}
	if newBookRecord.AuthorName != "" {
		requestedBook.AuthorName = newBookRecord.AuthorName
	}
	if newBookRecord.SerialNumber != "" {
		requestedBook.SerialNumber = newBookRecord.SerialNumber
	}
	// save the changes
	db.Save(&requestedBook)
	// marshall the new content
	response, _ := json.Marshal(requestedBook)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
