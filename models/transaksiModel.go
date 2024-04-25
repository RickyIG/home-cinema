package models

import "time"

type TransactionData struct {
	IDTransaksi int       `db:"id_transaksi"` // Use int for SERIAL data type
	IDUser      int       `db:"id_user"`      // int for INTEGER
	IDJadwal    string    `db:"id_jadwal"`    // String for VARCHAR
	TotalBayar  int       `db:"total_bayar"`  // int for INTEGER
	CreatedAt   time.Time `db:"created_at"`   // time.Time for TIMESTAMP'
	Tickets     []TicketData
}
