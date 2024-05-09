package routes

import (
	"github.com/SHREYANSHSAW/sawgo1/book store mangment/pkg/controllers"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("post")
	router.HandleFunc("/book/", controllers.GetBook).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("get")
	router.HandleFunc("book/{bookId}", controllers.UpdateBookById).Methods("update")
	router.HanleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("Delete")
}
