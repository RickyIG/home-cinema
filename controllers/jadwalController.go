package controllers

import (
	"fmt"
	"home-cinema/database"
	"home-cinema/models"
	"home-cinema/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllJadwal(c *gin.Context) {
	var (
		result gin.H
	)

	jadwals, err := repository.GetAllJadwal(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": jadwals,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertJadwal(c *gin.Context) {
	// Baca parameter dari request body
	var jadwal models.Jadwal
	err := c.ShouldBindJSON(&jadwal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buat string kunci unik
	jadwal.JadwalID = fmt.Sprintf("%d-%d-%d-%d-%d-%d-%d", jadwal.TanggalTayang.Year(), jadwal.TanggalTayang.Month(), jadwal.TanggalTayang.Day(), jadwal.JamTayang.Hour(), jadwal.JamTayang.Minute(), jadwal.FilmID, jadwal.StudioID)

	err = repository.InsertJadwal(database.DbConnection, jadwal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Jadwal",
	})

}

func UpdateJadwal(c *gin.Context) {
	var jadwal models.Jadwal
	id := c.Param("id_jadwal")

	err := c.ShouldBindJSON(&jadwal)
	if err != nil {
		panic(err)
	}

	jadwal.JadwalID = id

	err = repository.UpdateJadwal(database.DbConnection, jadwal)
	if err != nil {
		// panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Jadwal",
	})
}

func DeleteJadwal(c *gin.Context) {
	var jadwal models.Jadwal
	id := c.Param("id_jadwal")
	jadwal.JadwalID = id

	err := repository.DeleteJadwal(database.DbConnection, jadwal)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Jadwal",
	})
}

func GetJadwalByID(c *gin.Context) {
	var (
		result gin.H
	)

	jadwalIDParam := c.Param("id_jadwal")

	films, err := repository.GetJadwalByID(database.DbConnection, jadwalIDParam)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"result": films,
	}
	c.JSON(http.StatusOK, result)
}
