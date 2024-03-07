package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Corray333/peppe-store/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Product struct {
	ID       int    `json:"id" db:"product_id"`
	Type     string `json:"type" db:"type"`
	Name     string `json:"name" db:"name"`
	Quantity int    `json:"quantity" db:"quantity"`
	Price    int    `json:"price" db:"price"`
	Color    string `json:"color" db:"color"`
	Size     string `json:"size" db:"size"`
}

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

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		rows, err := storage.DB.Queryx("select products.product_id, size, quantity, price, color, type, name from product_size_color inner join products on products.product_id = product_size_color.product_id inner join (select  products.product_id, min(price) as min_price from product_size_color inner join products on products.product_id = product_size_color.product_id group by products.product_id order by min_price limit 50) min_prices on min_prices.product_id = product_size_color.product_id order by min_price, products.product_id;")
		if err != nil {
			log.Fatal(err)
		}

		product := Product{}
		products := []Product{}
		for rows.Next() {
			err := rows.StructScan(&product)
			if err != nil {
				log.Fatal(err)
			}
			products = append(products, product)
		}
		result := [][]Product{
			{products[0]},
		}
		counter := 0
		for i := 1; i < len(products); i++ {
			if products[i].Name == result[counter][0].Name {
				result[counter] = append(result[counter], products[i])
			} else {
				result = append(result, []Product{products[i]})
				counter++
			}
		}
		json.NewEncoder(w).Encode(result)
	})
	router.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		rows, err := storage.DB.Queryx("select products.product_id, size, quantity, price, color, type, name from product_size_color inner join products on products.product_id = product_size_color.product_id where products.product_id = " + chi.URLParam(r, "id") + ";")
		if err != nil {
			log.Fatal(err)
		}

		product := Product{}
		products := []Product{}
		for rows.Next() {
			err := rows.StructScan(&product)
			if err != nil {
				log.Fatal(err)
			}
			products = append(products, product)
		}
		json.NewEncoder(w).Encode(products)
	})
	router.Get("/products/featured", func(w http.ResponseWriter, r *http.Request) {
		rows, err := storage.DB.Queryx("select products.product_id, size, quantity, price, color, type, name from product_size_color inner join products on products.product_id = product_size_color.product_id where featured order by product_id;")
		if err != nil {
			log.Fatal(err)
		}

		product := Product{}
		products := []Product{}
		for rows.Next() {
			err := rows.StructScan(&product)
			if err != nil {
				log.Fatal(err)
			}
			products = append(products, product)
		}
		result := [][]Product{
			{products[0]},
		}
		counter := 0
		for i := 1; i < len(products); i++ {
			if products[i].Name == result[counter][0].Name {
				result[counter] = append(result[counter], products[i])
			} else {
				result = append(result, []Product{products[i]})
				counter++
			}
		}
		json.NewEncoder(w).Encode(result)
	})
	return router
}
