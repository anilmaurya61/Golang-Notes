package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Define a route for GET method
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World! This is a GET request.")
    }).Methods("GET")

    // Define a route for POST method
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "This is a POST request.")
    }).Methods("POST")

    // Define a route for PUT method
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "This is a PUT request.")
    }).Methods("PUT")

    // Define a route for DELETE method
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "This is a DELETE request.")
    }).Methods("DELETE")

    // Start the server on port 8080 using Gorilla Mux router
    fmt.Println("Server is running on http://localhost:8080")
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
