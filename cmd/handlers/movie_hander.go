package handlers

import "net/http"

// MovieCtx is blah blah
func MovieCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetMovies retrives all movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Movies"))
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Movie"))
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Movie"))
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Movie"))
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Movie"))
}
