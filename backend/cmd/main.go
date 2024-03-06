package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Product struct {
	ID       int      `json:"id"`
	Type     string   `json:"type"`
	Name     string   `json:"name"`
	Quantity int      `json:"quantity"`
	Price    int      `json:"price"`
	Img      string   `json:"img"`
	Colors   []string `json:"colors"`
}

func main() {
	products := []Product{
		{
			ID:       1,
			Type:     "Hoody",
			Name:     "Help for drug addicts",
			Quantity: 10,
			Price:    2990,
			Img:      "4035147859",
			Colors:   []string{"black"},
		},
		{
			ID:       2,
			Type:     "Hoody",
			Name:     "Bubblegum Crisis",
			Quantity: 13,
			Price:    2790,
			Img:      "4038595373",
			Colors:   []string{"black", "Pink"},
		},
		{
			ID:       3,
			Type:     "T-Shirt",
			Name:     "For sale",
			Quantity: 15,
			Price:    1990,
			Img:      "4037647308",
			Colors:   []string{"black", "white"},
		},
		{
			ID:       4,
			Type:     "Band",
			Name:     "Zamaj",
			Quantity: 100,
			Price:    990,
			Img:      "4128956269",
			Colors:   []string{"black"},
		},
		{
			ID:       5,
			Type:     "Longsleeve",
			Name:     "Sober",
			Quantity: 100,
			Price:    990,
			Img:      "4145732958",
			Colors:   []string{"white"},
		},
	}
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

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(products)
	})
	fmt.Println("Server is starting at: 127.0.0.1:3001 ...")
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}
