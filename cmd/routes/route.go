package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func NewRouter(port string) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userRouter(r)
	reviewRouter(r)
	movieRouter(r)

	log.Println("Server is running on port ", port)

	return r
}

func userRouter(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Users"))
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("User"))
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Create User"))
		})
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Update User"))
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Delete User"))
		})
	})
}

func reviewRouter(r *chi.Mux) {
	r.Route("/reviews", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Reviews"))
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Review"))
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Create Review"))
		})
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Update Review"))
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Delete Review"))
		})
	})
}

func movieRouter(r *chi.Mux) {
	r.Route("/movies", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Movies"))
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Movie"))
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Create Movie"))
		})
		r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Update Movie"))
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Delete Movie"))
		})
	})
}
