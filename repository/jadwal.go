package repository

import (
	"database/sql"
	"fmt"
	"home-cinema/models"
	"strconv"
)

func GetAllJadwal(db *sql.DB) (results []models.Jadwal, err error) {
	sql := "SELECT * FROM jadwals"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var jadwal = models.Jadwal{}

		err = rows.Scan(&jadwal.JadwalID, &jadwal.FilmID, &jadwal.StudioID, &jadwal.TanggalTayang, &jadwal.JamTayang, &jadwal.HargaTiket)
		if err != nil {
			panic(err)
		}

		results = append(results, jadwal)
	}

	return
}

func InsertJadwal(db *sql.DB, jadwal models.Jadwal) (err error) {
	sql := "INSERT INTO jadwals (id_jadwal, id_film, id_studio, tanggal_tayang, jam_tayang, harga_tiket) VALUES ($1, $2, $3, $4, $5, $6)"

	// year, month, day := jadwal.TanggalTayang.Date()
	// formattedDate := time.Date(year, month, day, 0, 0, 0, 0, jadwal.TanggalTayang.Location()).Format("2006-01-02T00:00:00Z")

	// formattedTime := jadwal.JamTayang.Format("2006-01-02T00:00:00Z")

	errs := db.QueryRow(sql, jadwal.JadwalID, jadwal.FilmID, jadwal.StudioID, jadwal.TanggalTayang, jadwal.JamTayang, jadwal.HargaTiket)

	// Get studio seat count from studios table
	var studioSeatCount int
	err = db.QueryRow(`SELECT jumlah_kursi FROM studios WHERE id_studio = $1`, jadwal.StudioID).Scan(&studioSeatCount)
	if err != nil {
		return err
	}

	// Generate and insert kursi records
	kursiID := 1

	for row := 1; row <= studioSeatCount/4; row++ {

		for col := 1; col <= 4; col++ {
			kursiName := string(rune('A'+row-1)) + "-" + strconv.Itoa(col)
			fmt.Println(kursiName)
			fmt.Println(row - 1)

			status := "tersedia"

			_, err = db.Exec(`INSERT INTO kursis (id_studio, nomor_kursi, status) VALUES ($1, $2, $3)`, jadwal.StudioID, kursiName, status)
			if err != nil {
				return err
			}

			kursiID++
		}
	}

	return errs.Err()
}

func UpdateJadwal(db *sql.DB, jadwal models.Jadwal) (err error) {
	sql := "UPDATE jadwals SET id_film = $1, id_studio = $2, tanggal_tayang = $3, jam_tayang = $4, harga_tiket = $5 WHERE id_jadwal = $6"

	// formattedTime := film.JamTayang.Format("2006-01-02T00:00:00Z")

	errs := db.QueryRow(sql, jadwal.FilmID, jadwal.StudioID, jadwal.TanggalTayang, jadwal.JamTayang, jadwal.HargaTiket, jadwal.JadwalID)

	return errs.Err()
}

func DeleteJadwal(db *sql.DB, jadwal models.Jadwal) (err error) {
	sql := "DELETE FROM jadwals WHERE id_jadwal = $1"

	errs := db.QueryRow(sql, jadwal.JadwalID)

	return errs.Err()
}

func GetJadwalByID(db *sql.DB, jadwalID string) (results []models.Jadwal, err error) {
	sql := "SELECT * FROM jadwals WHERE id_jadwal = $1"

	rows, err := db.Query(sql, jadwalID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var jadwal = models.Jadwal{}

		err = rows.Scan(&jadwal.JadwalID, &jadwal.FilmID, &jadwal.StudioID, &jadwal.TanggalTayang, &jadwal.JamTayang, &jadwal.HargaTiket)
		if err != nil {
			panic(err)
		}

		results = append(results, jadwal)
	}

	return
}
