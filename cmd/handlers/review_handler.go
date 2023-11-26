package handlers

import "net/http"

// ReviewCtx is blah
func ReviewCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

// GetReviews is
func GetReviews(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reviews"))
}

// GetReview is
func GetReview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Review"))
}

// CreateReview is
func CreateReview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Review"))
}

// UpdateReview is
func UpdateReview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Review"))
}

// DeleteReview is
func DeleteReview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Review"))
}
