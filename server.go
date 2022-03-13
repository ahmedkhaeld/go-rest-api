package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintln(resp, "Up and running...")
		if err != nil {
			return
		}
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server listening on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		return
	}

}
