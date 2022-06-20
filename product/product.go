package product

import (
	"Go_REST_API/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	ProdId    int    `json:"prod_id"`
	ProdName  string `json:"prod_name"`
	ProdPrice int    `json:"prod_price"`
	ProdCode  string `json:"prod_code"`
}

type JsonResponse struct {
	Type    string    `json:"type"`
	Data    []Product `json:"data"`
	Message string    `json:"message"`
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	//print message
	fmt.Println("Getting Product data")

	rows, err := db.Query("SELECT prod_id , prod_name , prod_price,  prod_code FROM product")

	config.CheckErr(err)

	var prodData []Product

	for rows.Next() {
		var prod_id int
		var prod_name string
		var prod_price int
		var prod_code string
		err = rows.Scan(&prod_id, &prod_name, &prod_price, &prod_code)
		prodData = append(prodData, Product{ProdId: prod_id, ProdName: prod_name, ProdPrice: prod_price, ProdCode: prod_code})
	}
	w.Header().Set("Content-Type", "application/json")
	var response = JsonResponse{Type: "success", Data: prodData, Message: "Fetch All Data complete"}

	json.NewEncoder(w).Encode(response)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	prod_name := r.FormValue("prod_name")
	prod_price := r.FormValue("prod_price")
	prod_code := r.FormValue("prod_code")
	pc_id := 1
	var lastID int

	fmt.Println(prod_name)
	fmt.Println(prod_price)
	fmt.Println(prod_code)

	insertQuery := "INSERT INTO product(prod_name , prod_price , prod_code , pc_id) VALUES ($1, $2 , $3 , $4) returning prod_id"

	err := db.QueryRow(insertQuery, prod_name, prod_price, prod_code, pc_id).Scan(&lastID)

	if err != nil {
		log.Fatalf("error query", err)
	}

	response := JsonResponse{Type: "success", Message: "Sukses Input data"}

	json.NewEncoder(w).Encode(response)

}
