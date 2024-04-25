package models

import "time"

type Film struct {
	IDFilm    int    `json:"id_film"`
	JudulFilm string `json:"judul_film"`
	Genre     string `json:"genre"`
	Sinopsis  string `json:"sinopsis"`
	Durasi    int    `json:"durasi"`
	// TanggalTayang time.Time `json:"tanggal_tayang"`
	// JamTayang     time.Time `json:"jam_tayang"`
	Rating    string    `json:"rating"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
