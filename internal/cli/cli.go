package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCLi(service *service.BookService) *BookCLI {
	return &BookCLI{service: service}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gobooks <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gobooks search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gobooks simulate <book_id> <book_id> <book_id> <book_id>")
			return
		}

		bookIds := os.Args[2:]
		cli.simulateReading(bookIds)
	}

}

func (cli *BookCLI) simulateReading(bookIdsStr []string) {
	var bookIds []int
	for _, idStr := range bookIdsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid book id:", idStr)
			continue
		}
		bookIds = append(bookIds, id)
	}

	responses := cli.service.SimulateMultipleReadings(bookIds, 2*time.Second)
	for _, response := range responses {
		fmt.Println(response)
	}
}

func (cli *BookCLI) searchBooks(bookName string) {
	books, err := cli.service.GetBooksByName(bookName)
	if err != nil {
		fmt.Println("Error searching books:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}

	fmt.Println(len(books), "books found:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Genre: %s\n", book.ID, book.Title, book.Author, book.Genre)
	}
}
