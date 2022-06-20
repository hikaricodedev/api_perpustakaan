package main

import (
	"Go_REST_API/product"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/product/", product.GetProduct).Methods("GET")
	router.HandleFunc("/product/store", product.InsertProduct).Methods("POST")
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
