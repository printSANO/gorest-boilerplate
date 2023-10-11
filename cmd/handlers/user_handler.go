package handlers

import "net/http"

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
