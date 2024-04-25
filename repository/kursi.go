package repository

import (
	"database/sql"
	"home-cinema/models"
)

func GetAllKursisByStudioID(db *sql.DB, studioID int) (results []models.Kursi, err error) {
	sql := "SELECT * FROM kursis WHERE id_studio = $1"

	rows, err := db.Query(sql, studioID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var kursi = models.Kursi{}

		err = rows.Scan(&kursi.IDKursi, &kursi.IDStudio, &kursi.NomorKursi, &kursi.Status)
		if err != nil {
			panic(err)
		}

		results = append(results, kursi)
	}

	return
}

func GetSpecifiedKursisByStudioID(db *sql.DB, studioID int, kursiID int) (results []models.Kursi, err error) {
	sql := "SELECT * FROM kursis WHERE id_studio = $1 AND id_kursi = $2"

	rows, err := db.Query(sql, studioID, kursiID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var kursi = models.Kursi{}

		err = rows.Scan(&kursi.IDKursi, &kursi.IDStudio, &kursi.NomorKursi, &kursi.Status)
		if err != nil {
			panic(err)
		}

		results = append(results, kursi)
	}

	return
}

func UpdateSpecifiedKursisByStudioID(db *sql.DB, kursi models.Kursi, studioID int, kursiID int) (err error) {
	sql := "UPDATE kursis SET nomor_kursi = $1, status = $2 WHERE id_studio = $3 AND id_kursi = $4"

	errs := db.QueryRow(sql, kursi.NomorKursi, kursi.Status, studioID, kursiID)

	return errs.Err()
}

func InsertKursiByStudioID(db *sql.DB, kursi models.Kursi) (err error) {
	sql := "INSERT INTO kursis (nomor_kursi, id_studio) VALUES ($1, $2)"

	errs := db.QueryRow(sql, kursi.NomorKursi, kursi.IDStudio)

	return errs.Err()
}

func DeleteKursiByStudioID(db *sql.DB, kursi models.Kursi) (err error) {
	sql := "DELETE FROM kursis WHERE id_kursi = $1 AND id_studio = $2"

	errs := db.QueryRow(sql, kursi.IDKursi, kursi.IDStudio)

	return errs.Err()
}
