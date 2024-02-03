package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jaipalsinghkhangarot/bookstore/pkg/models"
	"github.com/jaipalsinghkhangarot/bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/book/" {
		http.Error(response, "404 Not Found", http.StatusNotFound)
	}
	if request.Method != "GET" {
		http.Error(response, "method not supported", http.StatusNotFound)
	}
	newBooks := models.GetAllBook()
	res, _ := json.Marshal(newBooks)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(response, "method not supported", http.StatusNotFound)
	}
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "method not supported", http.StatusNotFound)
	}
	CreateBook := &models.Book{}
	utils.ParseBody(request, CreateBook)
	book := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func DeleteBook(response http.ResponseWriter, request *http.Request) {
	if request.Method != "DELETE" {
		http.Error(response, "method not supported", http.StatusNotFound)
	}
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {
	if request.Method != "PUT" {
		http.Error(response, "method not supported", http.StatusNotFound)
	}
	var updateBook = &models.Book{}
	utils.ParseBody(request, updateBook)
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booksDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	response.Header().Set("Content-Type", "pkglicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
