package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

/*
En GO, para manejar el request y el response se usa w y r
w es como va a ir el response. Sus headers, su contenido, etc.
r es como viene el request. Sus headers, su contenido, etc.
*/

// Ese metodo asigna de manera homologa los headers necesarios para que esta cosa funcione
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Este es un health check muy simple ademas de señalar los subdirectorios disponibles
func getRoot(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	io.WriteString(w, "This is root\n")
	io.WriteString(w, "Root may guide you to more useful places\n")
	io.WriteString(w, "/getusers/adduser/removeuser\n")
}

// Este metodo va a devolver en un JSON los usuarios y sus detalles.
func getData(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var err error
	var db *sql.DB

	url := "admin:admin@tcp(db:3306)/users"

	db, err = sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []map[string]interface{} //Crea una variable users que es una lista de "diccionarios"
	for rows.Next() {
		var id int
		var name string
		var email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("ID: %d, Name: %s\n", id, name) //%d de decimal (base 10, no que tenga punto), %s de string
		user := map[string]interface{}{
			"id":    id,
			"name":  name,
			"email": email,
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Este metodo va a agregar de un JSON el usuario requerido.
func appendData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var err error
	var db *sql.DB

	var user User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body. Must be a json with id, name, email.", http.StatusBadRequest)
		return
	}

	//io.WriteString(w, "Hello Maria!\n")
	url := "admin:admin@tcp(db:3306)/users"

	db, err = sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO users (id,name,email) VALUES (?, ?, ?);", user.ID, user.Name, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Regresa respuesta de exito
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}

// Este metodo va a eliminar un usuario
func dropData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var err error
	var db *sql.DB

	type ID struct {
		ID int `json:"id"`
	}

	var body ID

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body. Must be an id", http.StatusBadRequest)
		return
	}

	//io.WriteString(w, "Hello Maria!\n")
	url := "admin:admin@tcp(db:3306)/users"

	db, err = sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", body.ID)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Regresa respuesta de exito
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/getusers", getData)
	http.HandleFunc("/adduser", appendData)
	http.HandleFunc("/removeuser", dropData)

	http.ListenAndServe(":8000", nil)

}
