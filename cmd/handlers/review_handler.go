package handlers

import "net/http"

func ReviewCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func GetReviews(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reviews"))
}

func GetReview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Review"))
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Review"))
}

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Review"))
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Review"))
}
