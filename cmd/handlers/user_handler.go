package handlers

import (
	"log"
	"net/http"

	"github.com/printSANO/gorest-boilerplate/cmd/database"
	"github.com/printSANO/gorest-boilerplate/cmd/models"
)

// UserCtx is test
func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

// GetUsers is
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users"))
	query := "SELECT * FROM example2"

	rows, err := database.DBMain.Query(query)
	if err != nil {
		log.Println("Query Failed")
	}
	defer rows.Close()
	var results []models.Example2
	for rows.Next() {
		var result models.Example2
		err := rows.Scan(&result.ID, &result.Author, &result.Title, &result.CreatedAt)
		if err != nil {
			log.Println("Query Failed")
		}

		results = append(results, result)
	}
	log.Println(results)
}

// GetUser is
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User"))
}

// CreateUser is
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
	_, err := database.DBMain.Exec("INSERT INTO example2 (title, author) VALUES ($1, $2)", "Hi", "Hello")
	if err != nil {
		log.Printf("Data insertion failed. %v\n", err)
	} else {
		log.Println("Data inserted successfully.")
	}
}

// UpdateUser is
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

// DeleteUser is
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
