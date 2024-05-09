package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/saw/sawgo1/book store mangment/pkg/utils"
	"github.com/saw/sawgo1/book store management/pkg/models"
)

var NewBook models.Book

func getBook(w http.ResponseWriter. r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header.Set("content-Type", "pkglication/json")
	w.writeHeader(http.StatusOK)
	w.write(res)
}

func getBookById(w http.responsewriter, r *http.Request){
	vars := mux.vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")

	}
bookDetails, _ := models.GetBookById(ID)
res, _ := json.Marshal(bookDetails)
w.Header().Set("content-Type", "pkglication/json")
w.writeHeader(http.StatusOK)
w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.writer(res)
}

func DeleteBook(w http.ResponseWriter, r*http.Request){
	vars:= mux.vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().set("Content- Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) 
}

func UpdateBook(w http.ResponseWriter r *http.Request){
	var updateBook = &models.Book{}
	utils.Parse(r, updateBook)
	vars := mux.Vars(r)
	bookId : Vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	booksDetails, db:= models.GetBookById(ID)
	if updateBook.Name!=""{
		bookDetails.Name = updateBook.Name

	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication!= ""{
		bookDetails.Publication = updateBook.publication
	}
}