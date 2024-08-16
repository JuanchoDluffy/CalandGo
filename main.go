package main

import (
	"fmt"
	"log"
	"net/http"

	"CalandGo/routes"
)

func main() {
	//serve dummy page
	http.Handle("/", http.FileServer(http.Dir("./pages")))
	//define dummy route and useful one
	http.HandleFunc("/test", handler2)
	http.HandleFunc("/mats", routes.Handlesum)
	//print out to terminal for debug
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// dummy routes
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a response")
}
func handler2(w http.ResponseWriter, r *http.Request) {
	log.Println("got the test request")
	fmt.Fprint(w, "TEST")
}
