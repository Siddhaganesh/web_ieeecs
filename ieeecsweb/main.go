package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	dbConn, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	db = dbConn

	r := mux.NewRouter()

	r.HandleFunc("/books", getAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	rows, err := db.Query("SELECT id, title, author FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, b)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, _ := strconv.Atoi(params["id"])

	var b Book
	err := db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", bookID).Scan(&b.ID, &b.Title, &b.Author)
	if err != nil {
		log.Fatal(err)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	json.NewDecoder(r.Body).Decode(&b)

	_, err := db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", b.Title, b.Author)
	if err != nil {
		log.Fatal(err)
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, _ := strconv.Atoi(params["id"])

	var b Book
	json.NewDecoder(r.Body).Decode(&b)

	_, err := db.Exec("UPDATE books SET title = ?, author = ? WHERE id = ?", b.Title, b.Author, bookID)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, _ := strconv.Atoi(params["id"])

	_, err := db.Exec("DELETE FROM books WHERE id = ?", bookID)
	if err != nil {
		log.Fatal(err)
	}
}
