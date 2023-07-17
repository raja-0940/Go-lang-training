package main

import (
	"io"
	"net/http"

	//"log"
	"fmt"

	"encoding/json"

	"github.com/gorilla/mux"
	"mywebapp.com/DAL"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Write HAndler is called...")
	io.WriteString(w, "Welcome to GoLang Web Application!!")

}

func bookHandler(w http.ResponseWriter, r *http.Request) {

	//io.WriteString(w, "Welcome to GoLang Books Application!!")

	books := DAL.GetBooks() //it is returning [] books

	json.NewEncoder(w).Encode(books) // converting [] books to json & then sedning it as response

}

func bookbyIsbHandler(w http.ResponseWriter, r *http.Request) {

	//we have to retrive isbn value from request object

	var params = mux.Vars(r)

	var isbnValue = params["isbn"]

	var book = DAL.GetBookByIsbn(isbnValue)

	json.NewEncoder(w).Encode(book)

}

func bookPostHandler(w http.ResponseWriter, r *http.Request) {

	//io.WriteString(w, "Welcome to GoLang bookPostHandler Books Application!!")

	//r.Body : we will get data in the form of JSON

	var book DAL.Book

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&book) // converting json request body to book struct object

	fmt.Println(err)

	DAL.InsertBook(book)

}

func bookPutHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Welcome to GoLang bookPutHandler Books Application!!")

}

func bookDeleteHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	isbn := params["isbn"]

	DAL.DeleteBookByIsbn(isbn)

}
func main() {

	var router = mux.NewRouter()

	router.HandleFunc("/books", bookHandler).Methods("GET")
	router.HandleFunc("/books/{isbn}", bookbyIsbHandler).Methods("GET") //we r configuring route paramter

	router.HandleFunc("/books", bookPostHandler).Methods("POST")

	router.HandleFunc("/books", bookPutHandler).Methods("PUT")

	router.HandleFunc("/books", bookDeleteHandler).Methods("DELETE")

	//using HandleFunc we r configuring routing( mapping of incoming url with http handler)
	// http.HandleFunc("/", WelcomeHandler) // http://localhost:8000/

	// http.HandleFunc("/books", bookHandler) // http://localhost:8000/books

	http.ListenAndServe(":8000", router)

}
