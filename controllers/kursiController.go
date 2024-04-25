package controllers

import (
	"fmt"
	"home-cinema/database"
	"home-cinema/models"
	"home-cinema/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllKursisByStudioID(c *gin.Context) {
	var (
		result   gin.H
		studioID int
	)

	studioIDParam := c.Param("id_studio")
	studioID, err := strconv.Atoi(studioIDParam)
	if err != nil {
		result = gin.H{
			"error": "Invalid studio ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	kursis, err := repository.GetAllKursisByStudioID(database.DbConnection, studioID)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"result": kursis,
	}
	c.JSON(http.StatusOK, result)
}

func GetSpecifiedKursisByStudioID(c *gin.Context) {
	var (
		result   gin.H
		studioID int
		kursiID  int
	)

	studioIDParam := c.Param("id_studio")
	studioID, err := strconv.Atoi(studioIDParam)
	if err != nil {
		result = gin.H{
			"error": "Invalid studio ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	kursiIDParam := c.Param("id_kursi")
	kursiID, err = strconv.Atoi(kursiIDParam)
	if err != nil {
		result = gin.H{
			"error": "Invalid kursi ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	kursis, err := repository.GetSpecifiedKursisByStudioID(database.DbConnection, studioID, kursiID)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"result": kursis,
	}
	c.JSON(http.StatusOK, result)
}

func UpdateSpecifiedKursisByStudioID(c *gin.Context) {
	var (
		kursi    models.Kursi
		result   gin.H
		studioID int
		kursiID  int
	)

	studioIDParam := c.Param("id_studio")
	studioID, err := strconv.Atoi(studioIDParam)
	if err != nil {
		result = gin.H{
			"error": "Invalid studio ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	kursiIDParam := c.Param("id_kursi")
	kursiID, err = strconv.Atoi(kursiIDParam)
	if err != nil {
		result = gin.H{
			"error": "Invalid kursi ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	err = repository.UpdateSpecifiedKursisByStudioID(database.DbConnection, kursi, studioID, kursiID)
	if err != nil {
		// panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Kursi",
	})
}

func InsertKursiByStudioID(c *gin.Context) {
	var kursi models.Kursi
	// var result gin.H

	err := c.ShouldBindJSON(&kursi)
	if err != nil {
		panic(err)
	}

	studioIDParam := c.Param("id_studio")
	fmt.Println(studioIDParam)
	studioID, err := strconv.Atoi(studioIDParam)
	if err != nil {
		// result = gin.H{
		// 	"error": "Invalid studio ID",
		// }
		// c.JSON(http.StatusBadRequest, result)
		// return
		panic(err)
	}

	kursi.IDStudio = studioID

	err = repository.InsertKursiByStudioID(database.DbConnection, kursi)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Studio",
	})

}

func DeleteKursiByStudioID(c *gin.Context) {
	var kursi models.Kursi
	var result gin.H
	id, err := strconv.Atoi(c.Param("id_kursi"))
	if err != nil {
		panic(err)
	}

	kursi.IDKursi = int(id)

	studioIDParam := c.Param("id_studio")
	studioID, err := strconv.Atoi(studioIDParam)
	if err != nil {
		result = gin.H{
			"error": "Invalid studio ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	kursi.IDStudio = int(studioID)

	err = repository.DeleteKursiByStudioID(database.DbConnection, kursi)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Kursi",
	})
}
