package models

type Studio struct {
	IDStudio    int    `json:"id_studio"`
	NamaStudio  string `json:"nama_studio"`
	JumlahKursi int    `json:"jumlah_kursi"`
	TipeStudio  string `json:"tipe_studio"`
}
