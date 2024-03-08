package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Corray333/peppe-store/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int    `json:"id" db:"product_id"`
	Type        string `json:"type" db:"type"`
	Name        string `json:"name" db:"name"`
	Quantity    int    `json:"quantity" db:"quantity"`
	Price       int    `json:"price" db:"price"`
	Color       string `json:"color" db:"color"`
	Size        string `json:"size" db:"size"`
	Description string `json:"description" db:"description"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	minPrice := r.URL.Query().Get("minPrice")
	maxPrice := r.URL.Query().Get("maxPrice")
	searchQuery := r.URL.Query().Get("searchQuery")
	typeQuery := r.URL.Query().Get("typeQuery")
	colors := r.URL.Query()["color"]
	filterColors := "where color ="
	for _, color := range colors {
		filterColors += " '" + color + "' or color ="
	}
	if filterColors != "where color =" {
		filterColors = filterColors[:len(filterColors)-11]
	} else {
		filterColors = ""
	}
	filter1 := "where "
	if minPrice != "" {
		filter1 += "price >= " + minPrice + " and "
	}
	if maxPrice != "" {
		filter1 += "price <= " + maxPrice + " and "
	}
	if searchQuery != "" {
		filter1 += "name like '%" + searchQuery + "%' and "
	}
	if typeQuery != "" && typeQuery != "All" {
		filter1 += "type = '" + typeQuery + "' and "
	}
	if filter1 != "where " {
		filter1 = filter1[:len(filter1)-4]
	} else {
		filter1 = ""
	}
	fmt.Println("Filter colors: " + filterColors)
	fmt.Println("select products.product_id, size, quantity, price, color, type, name from product_size_color inner join products on products.product_id = product_size_color.product_id inner join (select  products.product_id, min(price) as min_price from product_size_color inner join products on products.product_id = product_size_color.product_id " + filterColors + " group by products.product_id order by min_price limit 50) min_prices on min_prices.product_id = product_size_color.product_id " + filter1 + "order by min_price, products.product_id;")
	rows, err := storage.DB.Queryx("select products.product_id, size, quantity, price, color, type, name from product_size_color inner join products on products.product_id = product_size_color.product_id inner join (select  products.product_id, min(price) as min_price from product_size_color inner join products on products.product_id = product_size_color.product_id " + filterColors + " group by products.product_id order by min_price limit 50) min_prices on min_prices.product_id = product_size_color.product_id " + filter1 + "order by min_price, products.product_id;")
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
		if products[i].Name == result[counter][0].Name && products[i].Type == result[counter][0].Type {
			result[counter] = append(result[counter], products[i])
		} else {
			result = append(result, []Product{products[i]})
			counter++
		}
	}
	json.NewEncoder(w).Encode(result)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	rows, err := storage.DB.Queryx("select products.product_id, size, quantity, price, color, type, name, description from product_size_color inner join products on products.product_id = product_size_color.product_id where products.product_id = " + chi.URLParam(r, "id") + ";")
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
}

func GetFeaturedProducts(w http.ResponseWriter, r *http.Request) {
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
}

func GetExactProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	color := r.URL.Query().Get("color")
	size := r.URL.Query().Get("size")

	rows, err := storage.DB.Queryx("select products.product_id, size, quantity, price, color, type, name from product_size_color inner join products on products.product_id = product_size_color.product_id where products.product_id = " + id + " and size = '" + size + "' and color='" + color + "';")
	if err != nil {
		log.Fatal(err)
	}

	product := Product{}
	for rows.Next() {
		err := rows.StructScan(&product)
		if err != nil {
			log.Fatal(err)
		}
	}
	json.NewEncoder(w).Encode(product)
}
