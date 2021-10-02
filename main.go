package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", home).Methods("GET")

	myRouter.HandleFunc("/posts", AllPosts).Methods("GET")
	myRouter.HandleFunc("/post/{title}/{author}", NewPost).Methods("POST")
	myRouter.HandleFunc("/post/{title}", DeletePost).Methods("DELETE")
	myRouter.HandleFunc("/post/{title}/{id}", UpdatePost).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	InitialMigration()

	handleRequests()
}
