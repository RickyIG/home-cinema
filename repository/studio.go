package repository

import (
	"database/sql"
	"home-cinema/models"
)

func GetAllStudio(db *sql.DB) (results []models.Studio, err error) {
	sql := "SELECT * FROM studios"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var studio = models.Studio{}

		err = rows.Scan(&studio.IDStudio, &studio.NamaStudio, &studio.JumlahKursi, &studio.TipeStudio)
		if err != nil {
			panic(err)
		}

		results = append(results, studio)
	}

	return
}

func InsertStudio(db *sql.DB, studio models.Studio) (err error) {
	sql := "INSERT INTO studios (nama_studio, jumlah_kursi, tipe_studio) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, studio.NamaStudio, studio.JumlahKursi, studio.TipeStudio)

	return errs.Err()
}

func UpdateStudio(db *sql.DB, studio models.Studio) (err error) {
	sql := "UPDATE studios SET nama_studio = $1, jumlah_kursi = $2, tipe_studio = $3 WHERE id_studio = $4"

	errs := db.QueryRow(sql, studio.NamaStudio, studio.JumlahKursi, studio.TipeStudio, studio.IDStudio)

	return errs.Err()
}

func DeleteStudio(db *sql.DB, studio models.Studio) (err error) {
	sql := "DELETE FROM studios WHERE id_studio = $1"

	errs := db.QueryRow(sql, studio.IDStudio)

	return errs.Err()
}

func GetStudioByID(db *sql.DB, categoryID int) (results []models.Studio, err error) {
	sql := "SELECT * FROM studios WHERE id_studio = $1"

	rows, err := db.Query(sql, categoryID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var studio = models.Studio{}

		err = rows.Scan(&studio.IDStudio, &studio.NamaStudio, &studio.JumlahKursi, &studio.TipeStudio)
		if err != nil {
			panic(err)
		}

		results = append(results, studio)
	}

	return
}

// func calculateJumlahKursi(namaStudio string) int {
// 	var jumlahKursi int
// 	err := db.QueryRow(`SELECT jumlah_kursi FROM studio_config WHERE nama_studio = $1`, namaStudio).Scan(&jumlahKursi)
// 	if err != nil {
// 		// Handle query error (studio might not exist)
// 		return 0
// 	}
// 	return jumlahKursi
// }
