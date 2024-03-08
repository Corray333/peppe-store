package storage

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() {
	var err error = nil
	DB, err = sqlx.Connect("postgres", os.Getenv("POSTGRES_CONNSTR"))
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
