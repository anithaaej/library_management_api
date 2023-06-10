package routers

import (
	"github.com/anithaa19/bms/books"
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/books", books.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/getbook/{id}", books.GetBookbyId).Methods("GET")
	router.HandleFunc("/api/book", books.CreateBook).Methods("POST")
	router.HandleFunc("/api/book/{id}", books.UpdateBookData).Methods("PUT")
	router.HandleFunc("/api/book/{id}", books.DeleteBookData).Methods("DELETE")
	router.HandleFunc("/api/deleteall", books.DeleteAllBookData).Methods("DELETE")

	return router
}
