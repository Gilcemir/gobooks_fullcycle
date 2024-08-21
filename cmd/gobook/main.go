package main

import (
	"database/sql"
	"gobooks/internal/cli"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}

	defer db.Close() //tudo daqui pra baixo é executado primeiro, depois esse trecho acontece

	bookService := service.NewBookService(db)
	bookHandlers := web.NewBookHandlers(bookService)

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		cli := cli.NewBookCLi(bookService)
		cli.Run()
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBookById)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)
	router.HandleFunc("GET /books/search", bookHandlers.SearchBooks)
	router.HandleFunc("POST /books/simulate", bookHandlers.SimulateMultipleReadings)

	http.ListenAndServe(":8080", router)
}
