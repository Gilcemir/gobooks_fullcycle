package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
	"strconv"
	"time"
)

type BookHandlers struct {
	service *service.BookService
}

func NewBookHandlers(service *service.BookService) *BookHandlers {
	return &BookHandlers{service: service}
}

func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()
	if err != nil {
		http.Error(w, "Failed to get Books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book service.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.CreateBook(&book)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) GetBookById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve book", http.StatusInternalServerError)
		return
	}

	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book service.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	book.ID = id

	if err := h.service.UpdateBook(&book); err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteBook(id); err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BookHandlers) SearchBooks(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if len(name) == 0 {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}

	books, err := h.service.GetBooksByName(name)
	if err != nil {
		http.Error(w, "Error retrieving books", http.StatusInternalServerError)
	}

	if books == nil || len(books) == 0 {
		http.Error(w, "Books not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandlers) SimulateMultipleReadings(w http.ResponseWriter, r *http.Request) {
	var bookIdsStr []string
	err := json.NewDecoder(r.Body).Decode(&bookIdsStr)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if len(bookIdsStr) == 0 {
		http.Error(w, "Book IDs cannot be empty", http.StatusBadRequest)
		return
	}

	var bookIds []int
	for _, idStr := range bookIdsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}
		bookIds = append(bookIds, id)
	}

	results := h.service.SimulateMultipleReadings(bookIds, 5*time.Second)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
