package main

import (
	"database/sql"
	"fmt"
	"home-cinema/controllers"
	"home-cinema/database"
	"home-cinema/middleware"
	"home-cinema/repository"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	//ENV Configuration
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// initiate gin
	router := gin.Default()

	// users
	router.POST("/api/register/admin", controllers.PostRegisterAdmin)
	router.POST("/api/register/user", controllers.PostRegisterUser)
	router.POST("/api/register/fintech", controllers.PostRegisterFintech)
	router.POST("/api/login", controllers.PostLogin)
	router.GET("/api/user/profile", middleware.JWTAuth(1), controllers.GetUserProfile)
	router.GET("/api/admin/profile", middleware.JWTAuth(2), controllers.GetUserProfile)
	router.GET("/api/fintech/profile", middleware.JWTAuth(3), controllers.GetUserProfile)

	// films
	router.GET("/films", controllers.GetAllFilm)
	router.GET("/films/:id", controllers.GetFilmByFilmID)
	router.POST("/films", middleware.JWTAuth(2), controllers.InsertFilm)
	router.PUT("/films/:id", middleware.JWTAuth(2), controllers.UpdateFilm)
	router.DELETE("/films/:id", middleware.JWTAuth(2), controllers.DeleteFilm)

	// studios
	router.GET("/studios", controllers.GetAllStudio)
	router.GET("/studios/:id_studio", controllers.GetStudioByStudioID)
	router.POST("/studios", middleware.JWTAuth(2), controllers.InsertStudio)
	router.PUT("/studios/:id_studio", middleware.JWTAuth(2), controllers.UpdateStudio)
	router.DELETE("/studios/:id_studio", middleware.JWTAuth(2), controllers.DeleteStudio)

	// kursis
	router.GET("/studios/:id_studio/seats", controllers.GetAllKursisByStudioID)
	router.GET("/studios/:id_studio/seats/:id_kursi", middleware.JWTAuth(1), controllers.GetSpecifiedKursisByStudioID)
	router.POST("/studios/:id_studio/seats", middleware.JWTAuth(2), controllers.InsertKursiByStudioID)
	router.PUT("/studios/:id_studio/seats/:id_kursi", middleware.JWTAuth(2), controllers.UpdateSpecifiedKursisByStudioID)
	router.DELETE("/studios/:id_studio/seats/:id_kursi", middleware.JWTAuth(2), controllers.DeleteKursiByStudioID)

	router.GET("/jadwal", controllers.GetAllJadwal)
	router.GET("/jadwal/:id_jadwal", middleware.JWTAuth(1), controllers.GetJadwalByID)
	router.POST("/jadwal", middleware.JWTAuth(2), controllers.InsertJadwal)
	router.PUT("/jadwal/:id_jadwal", middleware.JWTAuth(2), controllers.UpdateJadwal)
	router.DELETE("/jadwal/:id_jadwal", middleware.JWTAuth(2), controllers.DeleteJadwal)

	// transaksis
	// Create a TransaksiRepository instance
	transRepo := repository.NewTransaksiRepository(database.DbConnection)
	// Create a TransaksiController instance, passing the repository
	controller := controllers.NewTransaksiController(transRepo)
	// Define routes
	router.POST("/jadwal/:id_jadwal/transactions", middleware.JWTAuth(1), controller.CreateTransaction)
	// Riwayat Transaksi
	router.GET("/transactions", middleware.JWTAuth(1), controllers.GetUserTransactionHistory)
	// Cek Spesifik Riwayat Transaksi
	router.GET("/transactions/:id_transaksi", middleware.JWTAuth(1), controllers.GetUserTransactionHistoryByID)
	// Cek Spesifik Riwayat Transaksi beserta Detail
	router.GET("/transactions/:id_transaksi/details", middleware.JWTAuth(1), controllers.GetUserTransactionHistoryByIDWithDetails)

	// users - Top up by thirdparty
	router.PUT("/api/fintech/addbalance/user/:id_user", middleware.JWTAuth(3), controllers.UpdateUserBalance)

	router.GET("/transactions/tickets", middleware.JWTAuth(1), controllers.GetUserTicketHistory)
	router.GET("/transactions/:id_transaksi/tickets", middleware.JWTAuth(1), controllers.GetUserTicketHistoryByTransactionID)
	router.GET("/transactions/:id_transaksi/tickets/:id_ticket", middleware.JWTAuth(1), controllers.GetUserTicketHistoryByIDs)

	router.Run(":" + os.Getenv("PORT"))
}
