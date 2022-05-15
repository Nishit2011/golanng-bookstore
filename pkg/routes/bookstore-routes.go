package routes

import (
	"fmt"

	"github.com/Nishit2011/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	fmt.Printf("in routes")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")

	router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
}
