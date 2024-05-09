package models

import(
	"github.com/saw/gorm"
	"github.com/saw/sawgo1/book store management/PKG/config"

)

var db *gorm.DB

type Book struct{
	gorm.model
	name string 'gorm:""json:"name"'
	Author string 'json:"author"'
	publication string 'json:"publication'
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorn.db){
	var getBook Book
	db:= db.where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.where("ID=?", ID). Delete(book)
	return book
}