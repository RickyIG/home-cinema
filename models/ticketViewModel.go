package models

type TicketView struct {
	IDTicket     int    `db:"id_ticket"`     // Use int for SERIAL data type
	IDJadwal     string `db:"id_jadwal"`     // String for VARCHAR
	IDKursi      int    `db:"id_kursi"`      // int for INTEGER
	IDUser       int    `db:"id_user"`       // int for INTEGER
	IDTransaksi  int    `db:"id_transaksi"`  // int for INTEGER
	TicketStatus string `db:"ticket_status"` // String for VARCHAR(20)

}
