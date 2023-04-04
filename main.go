package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/favicon.ico", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode("favicon.ico")
	})

	router.HandleFunc("/{name}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		changeTest := "test 6:40"
		var message string
		if name == "" {
			message = "Hello World"
		} else {
			message = "Hello " + name + " " + changeTest
		}
		response := map[string]string{
			"message": message,
		}
		json.NewEncoder(rw).Encode(response)
	})

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "Welcome to Dockerized app",
		}
		json.NewEncoder(rw).Encode(response)
	})

	log.Println("Server is running!")
	fmt.Println(http.ListenAndServe(":8080", router))
}
