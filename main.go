package main

import (
	"encoding/json"
	"github.com/codeSmart2307/beginner_api/models"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Init books variable as a slice Book struct
var books []models.Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get request params
	// Lop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book) // Decodes request body
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - Not safe (could generate same ID more than once)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get request params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) // Similar to splice
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book) // Decodes request body
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get request params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) // Similar to splice
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Mux router
	r := mux.NewRouter()

	// Mock Data
	// TODO - Store data in DB
	books = append(books, models.Book{ID: "1", Isbn: "4483", Title: "Book One", Author: &models.Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, models.Book{ID: "2", Isbn: "7846", Title: "Book Two", Author: &models.Author{Firstname: "Jane", Lastname: "Austen"}})
	books = append(books, models.Book{ID: "3", Isbn: "2354", Title: "Book Three", Author: &models.Author{Firstname: "Steve", Lastname: "Smith"}})
	books = append(books, models.Book{ID: "4", Isbn: "8764", Title: "Book Four", Author: &models.Author{Firstname: "Emily", Lastname: "Blunt"}})

	// Route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
