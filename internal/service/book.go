package service

import (
	"database/sql"
	"fmt"
	"time"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(book *Book) error {
	query := "Insert into books (title, author, genre) values (?,?,?)"
	result, err := s.db.Exec(query, book.Title, book.Title, book.Genre)
	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	book.ID = int(lastInsertId)
	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "select id, title, author, genre from books"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBookById(id int) (*Book, error) {
	query := "select id, title, author, genre from books where id = ?"
	row := s.db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "update books set title = ?, author = ?, genre = ? where id = ?"
	_, err := s.db.Exec(query, book.Title, book.Title, book.Genre, book.ID)
	return err
}

func (s *BookService) DeleteBook(id int) error {
	query := "delete from books where id = ?"
	_, err := s.db.Exec(query, id)
	return err
}

func (s *BookService) SimulateReading(bookId int, duration time.Duration, results chan<- string) {
	book, err := s.GetBookById(bookId)
	if err != nil || book == nil {
		results <- fmt.Sprintf("Error reading book %d: %v", bookId, err)
		return
	}

	time.Sleep(duration)
	results <- fmt.Sprintf("Finished reading %s", book.Title)
}

func (s *BookService) SimulateMultipleReadings(bookIds []int, duration time.Duration) []string {
	booksCount := len(bookIds)
	results := make(chan string, booksCount)

	for _, id := range bookIds {
		go func(bookID int) {
			go s.SimulateReading(bookID, duration, results)
		}(id)
	}

	var responses []string
	for range bookIds {
		responses = append(responses, <-results)
	}

	close(results)
	return responses
}

func (s *BookService) GetBooksByName(name string) ([]Book, error) {
	query := "select id, title, author, genre from books where title like ?"
	rows, err := s.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
