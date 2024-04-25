package repository

import (
	"database/sql"
	"home-cinema/models"
	"time"
)

func GetAllFilm(db *sql.DB) (results []models.Film, err error) {
	sql := "SELECT * FROM films"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var film = models.Film{}

		err = rows.Scan(&film.IDFilm, &film.JudulFilm, &film.Genre, &film.Sinopsis, &film.Durasi, &film.Rating, &film.ImageURL, &film.CreatedAt, &film.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, film)
	}

	return
}

func InsertFilm(db *sql.DB, film models.Film) (err error) {
	sql := "INSERT INTO films (judul_film, genre, sinopsis, durasi, rating, image_url) VALUES ($1, $2, $3, $4, $5, $6)"

	// year, month, day := film.TanggalTayang.Date()
	// formattedDate := time.Date(year, month, day, 0, 0, 0, 0, film.TanggalTayang.Location()).Format("2006-01-02T00:00:00Z")

	// formattedTime := film.JamTayang.Format("2006-01-02T00:00:00Z")

	errs := db.QueryRow(sql, film.JudulFilm, film.Genre, film.Sinopsis, film.Durasi, film.Rating, film.ImageURL)

	return errs.Err()
}

func UpdateFilm(db *sql.DB, film models.Film) (err error) {
	sql := "UPDATE films SET judul_film = $1, genre = $2, sinopsis = $3, durasi = $4, rating = $5, image_url = $6, updated_at = $7 WHERE id_film = $8"

	// formattedTime := film.JamTayang.Format("2006-01-02T00:00:00Z")

	errs := db.QueryRow(sql, film.JudulFilm, film.Genre, film.Sinopsis, film.Durasi, film.Rating, film.ImageURL, time.Now(), film.IDFilm)

	return errs.Err()
}

func DeleteFilm(db *sql.DB, film models.Film) (err error) {
	sql := "DELETE FROM films WHERE id_film = $1"

	errs := db.QueryRow(sql, film.IDFilm)

	return errs.Err()
}

func GetFilmByID(db *sql.DB, categoryID int) (results []models.Film, err error) {
	sql := "SELECT * FROM films WHERE id_film = $1"

	rows, err := db.Query(sql, categoryID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var film = models.Film{}

		err = rows.Scan(&film.IDFilm, &film.JudulFilm, &film.Genre, &film.Sinopsis, &film.Durasi, &film.Rating, &film.ImageURL, &film.CreatedAt, &film.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, film)
	}

	return
}
