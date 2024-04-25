package repository

import (
	"database/sql"
	"errors"
	"time"

	"home-cinema/models"
)

type TransaksiRepository interface {
	BeginTransaction() (*sql.Tx, error)
	CreateTransaction(tx *sql.Tx, transactionData models.TransactionData, totalPrice int) (int, error)
	InsertTicket(tx *sql.Tx, transactionID int, ticketData *models.TicketData) error
	UpdateKursiStatus(tx *sql.Tx, kursiID int, status string) error
}

type transaksiRepositoryImpl struct {
	db *sql.DB
}

func NewTransaksiRepository(db *sql.DB) TransaksiRepository {
	return &transaksiRepositoryImpl{db: db}
}

func (repo *transaksiRepositoryImpl) BeginTransaction() (*sql.Tx, error) {
	return repo.db.Begin()
}

func (repo *transaksiRepositoryImpl) CreateTransaction(tx *sql.Tx, transactionData models.TransactionData, totalPrice int) (int, error) {
	// Insert transaction data into the 'transaksis' table
	insertQuery := `
    INSERT INTO transaksis (id_jadwal, id_user, created_at, total_bayar)
    VALUES ($1, $2, $3, $4)
    RETURNING id_transaksi
  `
	transactionData.TotalBayar = totalPrice
	transactionData.CreatedAt = time.Now()

	// fmt.Println("mau isi : ", transactionData)
	// fmt.Println("mau isi2  : ", transactionData.IDJadwal, transactionData.IDUser, transactionData.CreatedAt, transactionData.TotalBayar)

	rows, err := tx.Query(insertQuery, transactionData.IDJadwal, transactionData.IDUser, transactionData.CreatedAt, transactionData.TotalBayar)
	if err != nil {
		// panic(err)
		return 0, err
	}
	defer rows.Close()

	// Check if a row was actually inserted
	if !rows.Next() {
		return 0, errors.New("no rows affected") // Handle the case where no row was inserted
	}

	var transactionID int
	// fmt.Println("transactionID: ", transactionID)
	err = rows.Scan(&transactionID)
	if err != nil {
		return 0, err
	}

	return transactionID, nil

	// var transactionID int
	// err = rows.Scan(&transactionID)
	// if err != nil {
	// 	return 0, err
	// }

	// return transactionID, nil
}

func (repo *transaksiRepositoryImpl) InsertTicket(tx *sql.Tx, transactionID int, ticketData *models.TicketData) error {
	// Insert ticket data into the 'tickets' table, linking it to the transaction
	insertQuery := `
    INSERT INTO tickets (id_transaksi, id_jadwal, id_kursi, id_user, ticket_status)
    VALUES ($1, $2, $3, $4, $5)
  `

	ticketData.TicketStatus = "telah_dibayar"

	_, err := tx.Exec(insertQuery, transactionID, ticketData.IDJadwal, ticketData.IDKursi, ticketData.IDUser, ticketData.TicketStatus)
	if err != nil {
		return err
	}

	return nil
}

func (repo *transaksiRepositoryImpl) UpdateKursiStatus(tx *sql.Tx, kursiID int, status string) error {
	// Update kursi status to 'terisi'
	updateQuery := `
    UPDATE kursis
    SET status = $1
    WHERE id_kursi = $2
  `

	_, err := tx.Exec(updateQuery, status, kursiID)
	if err != nil {
		return err
	}

	return nil
}

func GetUserTransactionHistory(db *sql.DB, userID int) (results []models.TransactionView, err error) {
	sql := "SELECT id_transaksi, id_jadwal, id_user, total_bayar, created_at FROM transaksis WHERE id_user = $1"

	rows, err := db.Query(sql, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var transactionData = models.TransactionView{}

		err = rows.Scan(&transactionData.IDTransaksi, &transactionData.IDJadwal, &transactionData.IDUser, &transactionData.TotalBayar, &transactionData.CreatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, transactionData)
	}

	return
}

func GetUserTransactionHistoryByID(db *sql.DB, userID int, transactionID int) (results []models.TransactionView, err error) {
	sql := "SELECT id_transaksi, id_jadwal, id_user, total_bayar, created_at FROM transaksis WHERE id_user = $1 AND id_transaksi = $2"

	rows, err := db.Query(sql, userID, transactionID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var transactionData = models.TransactionView{}

		err = rows.Scan(&transactionData.IDTransaksi, &transactionData.IDJadwal, &transactionData.IDUser, &transactionData.TotalBayar, &transactionData.CreatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, transactionData)
	}

	return
}

// func GetUserTransactionHistoryByIDWithDetails(db *sql.DB, userID int, transaksiID int) (transactions []models.TransactionDetails, err error) {

// 	sql := `
//     SELECT t.id_transaksi, t.id_user, t.id_jadwal, t.total_bayar, t.created_at,
//            u.id_user AS user_id, u.username, u.email, u.first_name, u.last_name,
//            f.id_film AS film_id, f.judul_film, f.genre, f.sinopsis, f.durasi, f.rating, f.image_url,
//            s.id_studio AS studio_id, s.nama_studio, s.jumlah_kursi, s.tipe_studio,
//            j.id_jadwal AS jadwal_id, j.tanggal_tayang, j.jam_tayang, j.harga_tiket
//     FROM transaksis t
//     INNER JOIN users u ON t.id_user = u.id_user
//     INNER JOIN jadwals j ON t.id_jadwal = j.id_jadwal
//     INNER JOIN films f ON j.id_film = f.id_film
//     INNER JOIN studios s ON j.id_studio = s.id_studio
//     WHERE id_transaksi = $1 AND t.id_user = $2
//     ORDER BY t.created_at DESC;
//   `

// 	rows, err := db.Query(sql, userID, transaksiID)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(rows)

// 	for rows.Next() {
// 		var t models.TransactionDetails
// 		var u models.User
// 		var f models.Film
// 		var s models.Studio
// 		var j models.JadwalDetails

// 		err = rows.Scan(
// 			&t.IDTransaksi, &t.IDUser, &t.IDJadwal, &t.TotalBayar, &t.CreatedAt,
// 			&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName,
// 			&f.IDFilm, &f.JudulFilm, &f.Genre, &f.Sinopsis, &f.Durasi, &f.Rating, &f.ImageURL,
// 			&s.IDStudio, &s.NamaStudio, &s.JumlahKursi, &s.TipeStudio,
// 			&j.JadwalID, &j.TanggalTayang, &j.JamTayang, &j.HargaTiket,
// 		)
// 		if err != nil {
// 			panic(err)
// 		}

// 		t.User = u
// 		t.Jadwal.Film = f
// 		t.Jadwal.Studio = s
// 		t.Jadwal.JadwalID = j.JadwalID
// 		t.Jadwal.TanggalTayang = j.TanggalTayang
// 		t.Jadwal.JamTayang = j.JamTayang
// 		t.Jadwal.HargaTiket = j.HargaTiket

// 		transactions = append(transactions, t)
// 	}

// 	// formattedTime := film.JamTayang.Format("2006-01-02T00:00:00Z")
// 	fmt.Println(transactions)
// 	return transactions, err
// }

func GetUserTransactionHistoryByIDWithDetails(db *sql.DB, userID int, transaksiID int) ([]models.TransactionDetails, error) {

	sql := `
    SELECT
      t.id_transaksi,
      t.id_user,
      t.id_jadwal,
      t.total_bayar,
      t.created_at,
	  u.id_user,
      u.username,
      u.email,
      u.first_name,
      u.last_name,
	  u.password,
      u.phone_number, 
      u.balance,       
      u.id_role,        
      u.created_at AS user_created_at,  
      u.updated_at AS user_updated_at,  
      f.judul_film,
      f.genre,
      f.sinopsis,
      f.durasi,
      f.rating,
      f.image_url,
	  f.id_film,
	  f.created_at AS user_created_at,  
      f.updated_at AS user_updated_at,  
	  s.id_studio,
      s.nama_studio,
      s.jumlah_kursi,
      s.tipe_studio,
	  j.id_jadwal,
	  j.id_film,
	  j.id_studio,
      j.tanggal_tayang,
      j.jam_tayang,
      j.harga_tiket
	  FROM transaksis t
	  INNER JOIN users u ON t.id_user = u.id_user
	  INNER JOIN jadwals j ON t.id_jadwal = j.id_jadwal
	  INNER JOIN films f ON j.id_film = f.id_film
	  INNER JOIN studios s ON j.id_studio = s.id_studio
	  WHERE id_transaksi = $1 AND t.id_user = $2
	  ORDER BY t.created_at DESC;
	`

	rows, err := db.Query(sql, transaksiID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.TransactionDetails

	for rows.Next() {
		var t models.TransactionDetails
		err := rows.Scan(
			// Transaction fields
			&t.IDTransaksi,
			&t.IDUser,
			&t.IDJadwal,
			&t.TotalBayar,
			&t.CreatedAt,

			// User fields
			&t.User.UserID,
			&t.User.Username,
			&t.User.Email,
			&t.User.FirstName,
			&t.User.LastName,
			&t.User.Password,
			&t.User.PhoneNumber, // Added phone_number
			&t.User.Balance,     // Added balance
			&t.User.RoleID,      // Added id_role
			&t.User.CreatedAt,   // Aliased user_created_at
			&t.User.UpdatedAt,   // Aliased user_updated_at

			// Jadwal.Film fields
			&t.Jadwal.Film.JudulFilm,
			&t.Jadwal.Film.Genre,
			&t.Jadwal.Film.Sinopsis,
			&t.Jadwal.Film.Durasi,
			&t.Jadwal.Film.Rating,
			&t.Jadwal.Film.ImageURL,
			&t.Jadwal.Film.IDFilm,    // Added film_id
			&t.Jadwal.Film.CreatedAt, // Aliased film_created_at
			&t.Jadwal.Film.UpdatedAt, // Aliased film_updated_at

			// Jadwal.Studio fields
			&t.Jadwal.Studio.IDStudio,
			&t.Jadwal.Studio.NamaStudio,
			&t.Jadwal.Studio.JumlahKursi,
			&t.Jadwal.Studio.TipeStudio,

			// Jadwal fields
			&t.Jadwal.JadwalID,
			&t.Jadwal.FilmID,
			&t.Jadwal.StudioID,
			&t.Jadwal.TanggalTayang,
			&t.Jadwal.JamTayang,
			&t.Jadwal.HargaTiket,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, err
}
