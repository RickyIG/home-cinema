package models

import (
	"time"
)

type JadwalDetails struct {
	JadwalID      string    `json:"id_jadwal"`
	FilmID        int       `json:"id_film"`
	StudioID      int       `json:"id_studio"`
	TanggalTayang time.Time `json:"tanggal_tayang" time_format:"2006-01-02"`
	JamTayang     time.Time `json:"jam_tayang"`
	HargaTiket    int       `json:"harga_tiket"`

	Film   Film   `json:"film"`
	Studio Studio `json:"studio"`
}
