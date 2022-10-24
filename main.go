package main

import (
	"Go_REST_API/models"
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
	router.HandleFunc("/buku", models.GetBuku).Methods("GET")
	router.HandleFunc("/buku", models.InsertBuku).Methods("POST")
	router.HandleFunc("/buku/delete", models.DeleteBuku).Methods("POST")
	router.HandleFunc("/buku/update", models.UpdateBuku).Methods("POST")
	router.HandleFunc("/anggota", models.GetAnggota).Methods("GET")
	router.HandleFunc("/anggota", models.InsertAnggota).Methods("POST")
	router.HandleFunc("/anggota/delete", models.DeleteAnggota).Methods("POST")
	router.HandleFunc("/anggota/update", models.UpdateAnggota).Methods("POST")
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
