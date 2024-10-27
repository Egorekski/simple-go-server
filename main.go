package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		_, _ = io.WriteString(w, fmt.Sprintf("ParseForm() err: %s", err))
		return
	}

	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	_, err := io.WriteString(w, fmt.Sprintf("POST request successful\n"))
	if err != nil {
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")
	_, err = io.WriteString(w, fmt.Sprintf("Name: %s\nAdress: %s", name, address))
	if err != nil {
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	_, err := io.WriteString(w, "Hello!")
	if err != nil {
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
