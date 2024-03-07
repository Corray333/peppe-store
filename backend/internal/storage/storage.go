package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID       int      `json:"id" db:"product_id"`
	Type     string   `json:"type" db:"type"`
	Name     string   `json:"name" db:"name"`
	Quantity int      `json:"quantity" db:"quantity"`
	Price    int      `json:"price" db:"price"`
	Img      string   `json:"img" db:"img"`
	Colors   []string `json:"colors" db:"colors"`
	Sizes    []string `json:"sizes" db:"sizes"`
}

var DB *sqlx.DB

func ConnectDB() {
	var err error = nil
	DB, err = sqlx.Connect("postgres", "user=postgres dbname=Peppe sslmode=disable password=root host=localhost")
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection to the database
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
}

var Products = []Product{
	{
		ID:       1,
		Type:     "Hoody",
		Name:     "Help for drug addicts",
		Quantity: 10,
		Price:    2990,
		Colors:   []string{"black"},
		Sizes:    []string{"S", "M"},
	},
	{
		ID:       2,
		Type:     "Hoody",
		Name:     "Bubblegum Crisis",
		Quantity: 13,
		Price:    2790,
		Colors:   []string{"black", "pink"},
		Sizes:    []string{"S", "M", "L", "XL"},
	},
	{
		ID:       3,
		Type:     "T-Shirt",
		Name:     "For sale",
		Quantity: 15,
		Price:    1990,
		Img:      "4037647308",
		Colors:   []string{"black", "white"},
		Sizes:    []string{"S", "L", "XL"},
	},
	{
		ID:       4,
		Type:     "Band",
		Name:     "Zamaj",
		Quantity: 100,
		Price:    990,
		Img:      "4128956269",
		Colors:   []string{"black"},
		Sizes:    []string{"S", "M", "L"},
	},
	{
		ID:       5,
		Type:     "Longsleeve",
		Name:     "Sober",
		Quantity: 100,
		Price:    990,
		Img:      "4145732958",
		Colors:   []string{"white"},
		Sizes:    []string{"L", "XL"},
	},
}
