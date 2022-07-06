package models

import (
	"Go_REST_API/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Buku struct {
	IdBuku    int    `json:"id_buku"`
	KodeBuku  string `json:"kode_buku"`
	JudulBuku string `json:"nama_buku"`
	Deskripsi string `json:"deskripsi"`
	Tags      string `json:"tags"`
	Penerbit  string `json:"penerbit"`
	Pengarang string `json:"pengarang"`
	Editor    string `json:"editor"`
	Creator   string `json:"creator"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Buku `json:"data"`
	Message string `json:"message"`
}

func GetBuku(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	//print message
	fmt.Println("Getting Product data")

	rows, err := db.Query("SELECT * FROM buku")

	config.CheckErr(err)

	var dataBuku []Buku

	for rows.Next() {
		var id_buku int
		var kode_buku string
		var nama_buku string
		var deskripsi string
		var tags string
		var penerbit string
		var pengarang string
		var creator string
		var editor string

		err = rows.Scan(&id_buku, &kode_buku, &nama_buku, &deskripsi, &tags, &penerbit, &pengarang, &creator, &editor)
		dataBuku = append(dataBuku, Buku{IdBuku: id_buku, JudulBuku: nama_buku, Deskripsi: deskripsi, Tags: tags, Penerbit: penerbit, Pengarang: pengarang, Creator: creator, Editor: editor})
	}
	w.Header().Set("Content-Type", "application/json")
	var response = JsonResponse{Type: "success", Data: dataBuku, Message: "Fetch All Data complete"}

	json.NewEncoder(w).Encode(response)
}

func InsertBuku(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	nama_buku := r.FormValue("nama_buku")
	kode_buku := r.FormValue("kode_buku")
	deskripsi := r.FormValue("deskripsi")
	tags := r.FormValue("tags")
	penerbit := r.FormValue("penerbit")
	pengarang := r.FormValue("pengarang")
	creator := "user"
	editor := "user"
	var lastID int

	insertQuery := "INSERT INTO buku(nama_buku , kode_buku , deskripsi , tags , penerbit, pengarang,creator,editor) VALUES ($1, $2 , $3 , $4, $5 , $6 , $7 , $8) returning id_buku"

	err := db.QueryRow(insertQuery, nama_buku, kode_buku, deskripsi, tags, penerbit, pengarang, creator, editor).Scan(&lastID)

	if err != nil {
		log.Fatalf("error query", err)
	}

	response := JsonResponse{Type: "success", Message: "Sukses Input data"}

	json.NewEncoder(w).Encode(response)

}

func DeleteBuku(w http.ResponseWriter, r *http.Request) {

	id, _ := r.URL.Query()["id_buku"]
	id_buku := id[0]

	var response = JsonResponse{}

	fmt.Println(id_buku)

	if id_buku == "" {
		response = JsonResponse{Type: "error", Message: "Buku tidak ada"}
	} else {
		db := config.SetupDB()

		_, err := db.Query("DELETE FROM buku WHERE id_buku = $1", id_buku)

		config.CheckErr(err)

		response = JsonResponse{Type: "success", Message: "Buku berhasil di hapus"}
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateBuku(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	id, _ := r.URL.Query()["id_buku"]
	id_buku := id[0]

	nama_buku := r.FormValue("nama_buku")
	kode_buku := r.FormValue("kode_buku")
	deskripsi := r.FormValue("deskripsi")
	tags := r.FormValue("tags")
	penerbit := r.FormValue("penerbit")
	pengarang := r.FormValue("pengarang")
	editor := "user"
	var lastID int

	insertQuery := "UPDATE buku SET nama_buku = $1 , kode_buku = $2 , deskripsi = $3 , tags = $4 , penerbit = $5 , pengarang = $6 , editor = $7 WHERE id_buku = $8 returning id_buku "

	err := db.QueryRow(insertQuery, nama_buku, kode_buku, deskripsi, tags, penerbit, pengarang, editor, id_buku).Scan(&lastID)

	if err != nil {
		log.Fatalf("error query", err)
	}

	response := JsonResponse{Type: "success", Message: "Sukses Update data"}

	json.NewEncoder(w).Encode(response)

}
