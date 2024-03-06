package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	fmt.Println("Server is listening port 3000...")
	http.ListenAndServe(":3000", nil)
}
