package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Corray333/peppe-store/internal/config"
	"github.com/Corray333/peppe-store/internal/server"
	"github.com/Corray333/peppe-store/internal/storage"
	_ "github.com/lib/pq"
)

func main() {

	router := server.NewRouter()
	config.ConfigInit()
	storage.ConnectDB()

	fmt.Printf("Server is starting at: %s:%s ...", os.Getenv("APP_IP"), os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(os.Getenv("APP_IP")+":"+os.Getenv("APP_PORT"), router); err != nil {
		log.Fatal(err)
	}
}
