package controllers

import (
	"home-cinema/database"
	"home-cinema/models"
	"home-cinema/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStudio(c *gin.Context) {
	var (
		result gin.H
	)

	studios, err := repository.GetAllStudio(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": studios,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertStudio(c *gin.Context) {
	var studio models.Studio

	err := c.ShouldBindJSON(&studio)
	if err != nil {
		panic(err)
	}

	err = repository.InsertStudio(database.DbConnection, studio)
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

func UpdateStudio(c *gin.Context) {
	var studio models.Studio
	id, err := strconv.Atoi(c.Param("id_studio"))
	if err != nil {
		panic(err)
	}

	err = c.ShouldBindJSON(&studio)
	if err != nil {
		panic(err)
	}

	studio.IDStudio = int(id)

	err = repository.UpdateStudio(database.DbConnection, studio)

	if err != nil {
		// panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Studio",
	})
}

func DeleteStudio(c *gin.Context) {
	var studio models.Studio
	id, err := strconv.Atoi(c.Param("id_studio"))
	if err != nil {
		panic(err)
	}

	studio.IDStudio = int(id)

	err = repository.DeleteStudio(database.DbConnection, studio)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Studio",
	})
}

func GetStudioByStudioID(c *gin.Context) {
	var (
		result   gin.H
		studioID int
	)

	studioIDParam := c.Param("id_studio")
	studioID, err := strconv.Atoi(studioIDParam)
	if err != nil {
		result = gin.H{
			"error": "Invalid category ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	studios, err := repository.GetStudioByID(database.DbConnection, studioID)
	if err != nil {
		result = gin.H{
			"error": err,
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"result": studios,
	}
	c.JSON(http.StatusOK, result)
}
