package models

import (
	"Go_REST_API/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Anggota struct {
	IdAnggota int    `json:"id_amggota"`
	Na        string `json:"na"`
	Nama      string `json:"nama"`
	Alamat    string `json:"deskripsi"`
	Foto      string `json:"foto"`
	TglLahir  string `json:"tgl_lahir"`
	Editor    string `json:"editor"`
	Creator   string `json:"creator"`
}

type AnggotaResponse struct {
	Type    string    `json:"type"`
	Data    []Anggota `json:"data"`
	Message string    `json:"message"`
}

func GetAnggota(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	//print message
	fmt.Println("Getting Product data")

	rows, err := db.Query("SELECT * FROM anggota")

	config.CheckErr(err)

	var dataAnggota []Anggota

	for rows.Next() {
		var id_anggota int
		var na string
		var nama string
		var alamat string
		var foto string
		var tgl_lahir string
		var creator string
		var editor string

		err = rows.Scan(&id_anggota, &na, &nama, &alamat, foto, &tgl_lahir, &creator, &editor)
		dataAnggota = append(dataAnggota, Anggota{IdAnggota: id_anggota, Na: na, Nama: nama, Alamat: alamat, Foto: foto, Creator: creator, Editor: editor})
	}
	w.Header().Set("Content-Type", "application/json")
	var response = AnggotaResponse{Type: "success", Data: dataAnggota, Message: "Fetch All Data complete"}

	json.NewEncoder(w).Encode(response)
}

func InsertAnggota(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	na := r.FormValue("na")
	nama_anggota := r.FormValue("nama_anggota")
	alamat := r.FormValue("alamat")
	foto := r.FormValue("foto")
	tgl_lahir := r.FormValue("tgl_lahir")
	creator := "user"
	editor := "user"
	var lastID int

	insertQuery := "INSERT INTO anggota(na, nama_anggota, alamat, foto, tgl_lahir, creator, editor) VALUES ($1, $2 , $3 , $4, $5 , $6 , $7 , $8) returning id_buku"

	err := db.QueryRow(insertQuery, na, nama_anggota, alamat, foto, tgl_lahir, creator, editor).Scan(&lastID)

	if err != nil {
		log.Fatalf("error query", err)
	}

	response := JsonResponse{Type: "success", Message: "Sukses Input data"}

	json.NewEncoder(w).Encode(response)

}

func DeleteAnggota(w http.ResponseWriter, r *http.Request) {

	id, _ := r.URL.Query()["id_anggota"]
	id_anggota := id[0]

	var response = JsonResponse{}

	fmt.Println(id_anggota)

	if id_anggota == "" {
		response = JsonResponse{Type: "error", Message: "Anggota tidak ada"}
	} else {
		db := config.SetupDB()

		_, err := db.Query("DELETE FROM anggota WHERE id_anggota = $1", id_anggota)

		config.CheckErr(err)

		response = JsonResponse{Type: "success", Message: "Buku berhasil di hapus"}
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateAnggota(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	id, _ := r.URL.Query()["id_anggota"]
	id_anggota := id[0]

	na := r.FormValue("na")
	nama_anggota := r.FormValue("nama_anggota")
	alamat := r.FormValue("alamat")
	foto := r.FormValue("foto")
	tgl_lahir := r.FormValue("tgl_lahir")
	editor := "user"
	var lastID int

	insertQuery := "UPDATE anggota SET na = $1 , nama_anggota = $2 , alamat = $3 , foto = $4 , tgl_lahir = $5 , editor = $6 WHERE id_buku = $7 returning id_buku "

	err := db.QueryRow(insertQuery, na, nama_anggota, alamat, foto, tgl_lahir, editor, id_anggota).Scan(&lastID)

	if err != nil {
		log.Fatalf("error query", err)
	}

	response := JsonResponse{Type: "success", Message: "Sukses Update data"}

	json.NewEncoder(w).Encode(response)

}
