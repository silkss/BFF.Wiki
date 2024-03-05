package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT int = 8000
const HOST string = "localhost"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world!")
	})
	adress := fmt.Sprintf("%s:%d", HOST, PORT)
	log.Printf("starting server on %s\n", adress)
	http.ListenAndServe(adress, nil)
	log.Println("server closed.")
}
