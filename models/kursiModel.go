package models

type Kursi struct {
	IDKursi    int    `json:"id_kursi"`
	IDStudio   int    `json:"id_studio"`
	NomorKursi string `json:"nomor_kursi"`
	Status     string `json:"status"`
}
