package server

import (
	"github.com/Corray333/peppe-store/internal/server/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter() *chi.Mux {
	router := chi.NewMux()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/products", handlers.GetProducts)
	router.Get("/products/{id}", handlers.GetProduct)
	router.Get("/products/{id}/type", handlers.GetExactProduct)
	router.Get("/products/featured", handlers.GetFeaturedProducts)
	return router
}
