package main

import (
	"fmt"
	"log"
	"net/http"
	"postgress_go/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
