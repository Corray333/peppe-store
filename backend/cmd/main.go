package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Corray333/peppe-store/internal/server"
	"github.com/Corray333/peppe-store/internal/storage"
	_ "github.com/lib/pq"
)

func main() {

	router := server.NewRouter()
	storage.ConnectDB()

	fmt.Println("Server is starting at: 127.0.0.1:3001 ...")
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}
