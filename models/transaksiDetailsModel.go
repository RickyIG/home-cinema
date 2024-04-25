package models

import "time"

type TransactionDetails struct {
	IDTransaksi int       `json:"id_transaksi" db:"id_transaksi"` // Use int for SERIAL data type
	IDUser      int       `json:"id_user" db:"id_user"`           // int for INTEGER
	IDJadwal    string    `json:"id_jadwal" db:"id_jadwal"`       // String for VARCHAR
	TotalBayar  int       `json:"total_bayar" db:"total_bayar"`   // int for INTEGER
	CreatedAt   time.Time `json:"created_at" db:"created_at"`     // time.Time for TIMESTAMP'

	// User (nested struct)
	User User `json:"user"` // Assuming a separate User struct exists

	// Jadwal (nested struct)
	Jadwal JadwalDetails `json:"jadwal"` // Assuming a separate Jadwal struct exists

	// Film Film `json:"film"`

	// Studio Studio `json:"studio"`
}
