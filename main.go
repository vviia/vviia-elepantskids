package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vviia/elepantskids/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server started on: http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
