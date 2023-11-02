package routes

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/printSANO/gorest-boilerplate/cmd/handlers"
	"github.com/printSANO/gorest-boilerplate/config"
)

func NewRouter(port string, basePath string) *chi.Mux {
	r := chi.NewRouter()

	config.BasePath = basePath

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route(basePath, func(r chi.Router) {
		userRouter(r)
		reviewRouter(r)
		movieRouter(r)
	})

	log.Println("Server is running on port", port)

	return r
}

func userRouter(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Use(handlers.UserCtx)
		r.Get("/", handlers.GetUsers)
		r.Get("/{id}", handlers.GetUser)
		r.Post("/", handlers.CreateUser)
		r.Put("/{id}", handlers.UpdateUser)
		r.Delete("/{id}", handlers.DeleteUser)
	})
}

func reviewRouter(r chi.Router) {
	r.Route("/reviews", func(r chi.Router) {
		r.Use(handlers.ReviewCtx)
		r.Get("/", handlers.GetReviews)
		r.Get("/{id}", handlers.GetReview)
		r.Post("/", handlers.CreateReview)
		r.Put("/{id}", handlers.UpdateReview)
		r.Delete("/{id}", handlers.DeleteReview)
	})
}

func movieRouter(r chi.Router) {
	r.Route("/movies", func(r chi.Router) {
		r.Use(handlers.MovieCtx)
		r.Get("/", handlers.GetMovies)
		r.Get("/{id}", handlers.GetMovie)
		r.Post("/", handlers.CreateMovie)
		r.Put("/{id}", handlers.UpdateMovie)
		r.Delete("/{id}", handlers.DeleteMovie)
	})
}
