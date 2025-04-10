package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	io.WriteString(w, "This is root\n")
}

func getMaria(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	var err error
	var db *sql.DB
	io.WriteString(w, "Hello World!\n")
	dsn := "root:@tcp(localhost:3306)/users"

	db, err = sql.Open("maridb", dsn)

	// Example: Fetch data from a table
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}

func main() {
	http.HandleFunc("/", getRoot)

	http.ListenAndServe(":8000", nil)

}
