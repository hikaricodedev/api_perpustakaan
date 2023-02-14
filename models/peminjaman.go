package models

import (
	"Go_REST_API/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Peminjaman struct {
	IdPeminjaman int    `json:"id_peminjaman"`
	KdPeminjaman string `json:"kd_peminjaman"`
	IdAnggota    int    `json:"id_anggota"`
	IdBuku       int    `json:"id_buku"`
	TglPinjam    string `json:"tgl_pinjam"`
	TglKembali   string `json:"tgl_kembali"`
	Status       int    `json:"status"`
}

type PeminjamanResponse struct {
	Type    string       `json:"type"`
	Data    []Peminjaman `json:"data"`
	Message string       `json:"message"`
}

func GetPeminjaman(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	//print message
	fmt.Println("Getting Product data")

	rows, err := db.Query("SELECT * FROM buku")

	config.CheckErr(err)

	var dataPinjam []Peminjaman

	for rows.Next() {
		var id_peminjaman int
		var kd_peminjaman string
		var id_anggota int
		var id_buku int
		var tgl_pinjam string
		var tgl_kembali string
		var status int

		err = rows.Scan(&id_peminjaman, &kd_peminjaman, &id_anggota, &id_buku, &tgl_pinjam, &tgl_kembali, &status)
		dataPinjam = append(dataPinjam, Peminjaman{IdPeminjaman: id_peminjaman, KdPeminjaman: kd_peminjaman, IdBuku: id_buku, IdAnggota: id_anggota, TglPinjam: tgl_pinjam, TglKembali: tgl_kembali, Status: status})
	}
	w.Header().Set("Content-Type", "application/json")
	var response = PeminjamanResponse{Type: "success", Data: dataPinjam, Message: "Fetch All Data complete"}

	json.NewEncoder(w).Encode(response)
}

func InsertPeminjaman(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	var peminjamanNew Peminjaman
	errNew := json.NewDecoder(r.Body).Decode(&peminjamanNew)
	if errNew != nil {
		http.Error(w, errNew.Error(), http.StatusBadRequest)
	}

	kd_peminjaman := peminjamanNew.KdPeminjaman
	id_anggota := peminjamanNew.IdAnggota
	id_buku := peminjamanNew.IdBuku
	tgl_pinjam := peminjamanNew.TglPinjam
	tgl_kembali := peminjamanNew.TglKembali
	status := peminjamanNew.Status
	creator := "user"
	editor := "user"
	var lastID int

	insertQuery := "INSERT INTO peminjaman(kd_anggota , id_anggota , id_buku , tgl_pinjam , tgl_kembali, status,creator,editor) VALUES ($1, $2 , $3 , $4, $5 , $6 , $7 , $8) returning id_peminjaman"

	err := db.QueryRow(insertQuery, kd_peminjaman, id_anggota, id_buku, tgl_pinjam, tgl_kembali, status, creator, editor).Scan(&lastID)

	if err != nil {
		log.Fatalf("error query", err)
	}

	response := JsonResponse{Type: "success", Message: "Sukses Input data"}

	json.NewEncoder(w).Encode(response)

}

func DeletePeminjaman(w http.ResponseWriter, r *http.Request) {

	id, _ := r.URL.Query()["id_peminjaman"]
	id_peminjaman := id[0]

	var response = JsonResponse{}

	fmt.Println(id_peminjaman)

	if id_peminjaman == "" {
		response = JsonResponse{Type: "error", Message: "Buku tidak ada"}
	} else {
		db := config.SetupDB()

		_, err := db.Query("DELETE FROM peminjaman WHERE id_peminjaman = $1", id_peminjaman)

		config.CheckErr(err)

		response = JsonResponse{Type: "success", Message: "Buku berhasil di hapus"}
	}

	json.NewEncoder(w).Encode(response)
}

func UpdatePeminjaman(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDB()

	var peminjamanNew Peminjaman
	errNew := json.NewDecoder(r.Body).Decode(&peminjamanNew)
	if errNew != nil {
		http.Error(w, errNew.Error(), http.StatusBadRequest)
	}

	id, _ := r.URL.Query()["id_peminjaman"]
	id_peminjaman := id[0]

	kd_peminjaman := peminjamanNew.KdPeminjaman
	id_anggota := peminjamanNew.IdAnggota
	id_buku := peminjamanNew.IdBuku
	tgl_pinjam := peminjamanNew.TglPinjam
	tgl_kembali := peminjamanNew.TglKembali
	status := peminjamanNew.Status
	editor := "user"
	var lastID int

	insertQuery := "UPDATE buku SET kd_peminjaman = $1 , id_anggota = $2 , id_buku = $3 , tgl_pinjam = $4 , tgl_kembali = $5 , status = $6 , editor = $7 WHERE id_peminjaman = $8 returning id_peminjaman "

	err := db.QueryRow(insertQuery, kd_peminjaman, id_anggota, id_buku, tgl_pinjam, tgl_kembali, status, editor, id_peminjaman).Scan(&lastID)

	if err != nil {
		log.Fatalf("error query", err)
	}

	response := JsonResponse{Type: "success", Message: "Sukses Update data"}

	json.NewEncoder(w).Encode(response)

}
