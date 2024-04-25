package repository

import (
	"database/sql"

	"home-cinema/models"
)

func GetUserTicketHistory(db *sql.DB, userID int) (results []models.TicketView, err error) {
	sql := "SELECT id_ticket, id_jadwal, id_kursi, id_user, id_transaksi, ticket_status FROM tickets WHERE id_user = $1"

	rows, err := db.Query(sql, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ticketView = models.TicketView{}

		err = rows.Scan(&ticketView.IDTicket, &ticketView.IDJadwal, &ticketView.IDKursi, &ticketView.IDUser, &ticketView.IDTransaksi, &ticketView.TicketStatus)
		if err != nil {
			panic(err)
		}

		results = append(results, ticketView)
	}

	return
}

func GetUserTicketHistoryByTransactionID(db *sql.DB, userID int, transactionID int) (results []models.TicketView, err error) {
	sql := "SELECT id_ticket, id_jadwal, id_kursi, id_user, id_transaksi, ticket_status FROM tickets WHERE id_user = $1 AND id_transaksi = $2"

	rows, err := db.Query(sql, userID, transactionID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ticketView = models.TicketView{}

		err = rows.Scan(&ticketView.IDTicket, &ticketView.IDJadwal, &ticketView.IDKursi, &ticketView.IDUser, &ticketView.IDTransaksi, &ticketView.TicketStatus)
		if err != nil {
			panic(err)
		}

		results = append(results, ticketView)
	}

	return
}

func GetUserTicketHistoryByID(db *sql.DB, userID int, transactionID int, ticketID int) (results []models.TicketView, err error) {
	sql := "SELECT id_ticket, id_jadwal, id_kursi, id_user, id_transaksi, ticket_status FROM tickets WHERE id_user = $1 AND id_transaksi = $2 AND id_ticket = $3"

	rows, err := db.Query(sql, userID, transactionID, ticketID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ticketView = models.TicketView{}

		err = rows.Scan(&ticketView.IDTicket, &ticketView.IDJadwal, &ticketView.IDKursi, &ticketView.IDUser, &ticketView.IDTransaksi, &ticketView.TicketStatus)
		if err != nil {
			panic(err)
		}

		results = append(results, ticketView)
	}

	return
}
