package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	//io.WriteString(w, "Hello Maria!\n")
	dsn := "admin:admin@tcp(db:3306)/users"

	db, err = sql.Open("mysql", dsn)

	// Example: Fetch data from a table
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
		user := map[string]interface{}{
			"id":   id,
			"name": name,
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/data", getMaria)

	http.ListenAndServe(":8000", nil)

}
