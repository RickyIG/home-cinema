package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"home-cinema/database"
	"home-cinema/models"
	"home-cinema/repository"
)

type TransaksiController struct {
	transRepo repository.TransaksiRepository
	userRepo  repository.UserRepository
}

func NewTransaksiController(transRepo repository.TransaksiRepository) *TransaksiController {
	return &TransaksiController{transRepo: transRepo}
}

func (controller *TransaksiController) CreateTransaction(c *gin.Context) {
	var transactionData models.TransactionData

	transactionData.IDUser = c.GetInt("userID")
	transactionData.IDJadwal = c.Param("id_jadwal")

	err := c.BindJSON(&transactionData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for nomer := range transactionData.Tickets {
		transactionData.Tickets[nomer].IDUser = transactionData.IDUser
		transactionData.Tickets[nomer].IDJadwal = transactionData.IDJadwal
		transactionData.Tickets[nomer].TicketStatus = "belum_dibayar"
	}

	// fmt.Println(transactionData.Tickets)

	// Validate transaction data
	err = validateTransactionData(&transactionData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calculate total price
	totalPrice, err := calculateTotalPrice(transactionData.Tickets)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("total ", totalPrice)

	// Check if user balance is sufficient
	fmt.Println("iduser ", transactionData.IDUser)
	// Create UserRepository instance
	userRepository := repository.NewUserRepository(database.DbConnection)
	currentUserBalance, err := userRepository.GetBalance(transactionData.IDUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if currentUserBalance < totalPrice {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Saldo tidak mencukupi"})
		return
	}

	// Start transaction
	tx, err := controller.transRepo.BeginTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer tx.Rollback()

	// Create transaction record
	transactionID, err := controller.transRepo.CreateTransaction(tx, transactionData, totalPrice)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("cek")
	// Insert each ticket data and update kursi status
	for _, ticketData := range transactionData.Tickets {
		fmt.Println(ticketData)
		err := controller.transRepo.InsertTicket(tx, transactionID, &ticketData)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = controller.transRepo.UpdateKursiStatus(tx, ticketData.IDKursi, "terisi")
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Deduct transaction amount from user balance
	newBalance := currentUserBalance - totalPrice
	err = controller.userRepo.UpdateBalance(tx, transactionData.IDUser, newBalance)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Commit transaction if all steps are successful
	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaksi berhasil dibuat", "transaction_id": transactionID})
}

func calculateTotalPrice(tickets []models.TicketData) (int, error) {
	totalPrice := 0

	// Fetch schedule information for each ticket
	for _, ticket := range tickets {
		scheduleData, err := getScheduleData(database.DbConnection, ticket.IDJadwal, ticket.IDKursi)
		fmt.Println(scheduleData)
		if err != nil {
			return 0, err
		}

		ticket.HargaTiket = scheduleData.HargaTiket // Update ticket with price from schedule

		totalPrice += ticket.HargaTiket
	}

	// // Calculate total price using updated ticket prices
	// for _, ticket := range tickets {
	// 	totalPrice += ticket.HargaTiket
	// }

	return totalPrice, nil
}

func validateTransactionData(transactionData *models.TransactionData) error {
	// Validate transaction data
	if transactionData.IDJadwal == "" {
		return fmt.Errorf("ID jadwal tidak boleh kosong")
	}

	if transactionData.IDUser == 0 {
		return fmt.Errorf("ID user tidak boleh kosong")
	}

	// if transactionData.TotalBayar == 0 {
	// 	return fmt.Errorf("total bayar tidak boleh kosong")
	// }

	// Validate individual ticket data
	for _, ticketData := range transactionData.Tickets {
		if ticketData.IDKursi == 0 {
			return fmt.Errorf("id kursi pada tiket #%d tidak boleh kosong", ticketData.IDTicket)
		}

		if ticketData.IDUser == 0 {
			return fmt.Errorf("id user pada tiket #%d tidak boleh kosong", ticketData.IDTicket)
		}

		if ticketData.TicketStatus != "belum_dibayar" {
			return fmt.Errorf("status tiket pada tiket #%d harus 'belum_dibayar'", ticketData.IDKursi)
		}

		// Check if kursi exists and is available
		err := checkKursiAvailability(ticketData.IDKursi)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkKursiAvailability(kursiID int) error {
	// Check if kursi exists
	query := `
	  SELECT COUNT(*) FROM kursis
	  WHERE id_kursi = $1
	`

	var kursiCount int
	err := database.DbConnection.QueryRow(query, kursiID).Scan(&kursiCount)
	if err != nil {
		return fmt.Errorf("kursi dengan ID #%d tidak ditemukan", kursiID)
	}

	if kursiCount == 0 {
		return fmt.Errorf("kursi dengan ID #%d tidak tersedia", kursiID)
	}

	// Check if kursi is available
	query = `
	  SELECT status FROM kursis
	  WHERE id_kursi = $1
	`

	var kursiStatus string
	err = database.DbConnection.QueryRow(query, kursiID).Scan(&kursiStatus)
	if err != nil {
		return err
	}

	if kursiStatus != "tersedia" {
		return fmt.Errorf("kursi dengan ID #%d tidak tersedia", kursiID)
	}

	return nil
}

func getScheduleData(db *sql.DB, idJadwal string, idKursi int) (*models.Jadwal, error) {
	// Query the database to fetch schedule data
	query := `
	  SELECT j.id_jadwal, j.id_studio, j.harga_tiket
	  FROM jadwals j
	  INNER JOIN kursis k ON j.id_studio = k.id_studio
	  WHERE j.id_jadwal = $1::text AND k.id_kursi = $2
	`
	fmt.Println("test")
	fmt.Println(idJadwal, idKursi)

	rows, err := db.Query(query, idJadwal, idKursi)
	if err != nil {
		panic(err)
		// return nil, err
	}
	defer rows.Close()

	var scheduleData models.Jadwal
	if rows.Next() {
		err := rows.Scan(&scheduleData.JadwalID, &scheduleData.StudioID, &scheduleData.HargaTiket)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("schedule data not found for ID jadwal %s and ID kursi %d", idJadwal, idKursi)
	}

	fmt.Println("test: ", &scheduleData)

	return &scheduleData, nil
}

func GetUserTransactionHistory(c *gin.Context) {
	userID, isExist := c.Get("userID")
	if !isExist {
		c.JSON(http.StatusBadRequest, "User doesn't exist.")
		return
	}

	userTransactionHistory, err := repository.GetUserTransactionHistory(database.DbConnection, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := gin.H{
		"result": userTransactionHistory,
	}
	c.JSON(http.StatusOK, result)
}

func GetUserTransactionHistoryByID(c *gin.Context) {
	userID, isExist := c.Get("userID")
	if !isExist {
		c.JSON(http.StatusBadRequest, "User doesn't exist.")
		return
	}

	transaksiIDParam := c.Param("id_transaksi")
	transaksiID, err := strconv.Atoi(transaksiIDParam)
	if err != nil {
		result := gin.H{
			"error": "Invalid transaksi ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	userTransactionHistoryByID, err := repository.GetUserTransactionHistoryByID(database.DbConnection, userID.(int), transaksiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := gin.H{
		"result": userTransactionHistoryByID,
	}
	c.JSON(http.StatusOK, result)
}

func GetUserTransactionHistoryByIDWithDetails(c *gin.Context) {
	userID, isExist := c.Get("userID")
	if !isExist {
		c.JSON(http.StatusBadRequest, "User doesn't exist.")
		return
	}

	transaksiIDParam := c.Param("id_transaksi")
	transaksiID, err := strconv.Atoi(transaksiIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaksi ID"})
		return
	}

	transactions, err := repository.GetUserTransactionHistoryByIDWithDetails(database.DbConnection, userID.(int), transaksiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
