package controllers

import (
	"home-cinema/database"
	"home-cinema/models"
	"home-cinema/repository"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllFilm(c *gin.Context) {
	var (
		result gin.H
	)

	films, err := repository.GetAllFilm(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": films,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertFilm(c *gin.Context) {
	var film models.Film

	err := c.ShouldBindJSON(&film)
	if err != nil {
		panic(err)
	}

	// Validate image URL
	// image url yang valid, contoh: "https://www.example.com/images/photo.jpg"

	// imageURLRegex, err := regexp.Compile(`(http(s?):)([/|.|\w|\s|-])*\.(?:png|jpg|jpeg|gif|png|svg)`)
	imageURLRegex, err := regexp.Compile(`(http)?s?:?(\/\/[^"']*\.(?:png|jpg|jpeg|gif|png|svg))`)

	if err != nil {
		// panic(fmt.Errorf("error compiling regular expression: %w", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if !imageURLRegex.MatchString(film.ImageURL) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid image URL format",
		})
		return
	}

	err = repository.InsertFilm(database.DbConnection, film)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Film",
	})

}

func UpdateFilm(c *gin.Context) {
	var film models.Film
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = c.ShouldBindJSON(&film)
	if err != nil {
		panic(err)
	}

	film.IDFilm = int(id)

	// Validate image URL
	// image url yang valid, contoh: "https://www.example.com/images/photo.jpg"
	// imageURLRegex, err := regexp.Compile(`(http(s?):)([/|.|\w|\s|-])*\.(?:png|jpg|jpeg|gif|png|svg)`)
	imageURLRegex, err := regexp.Compile(`(http)?s?:?(\/\/[^"']*\.(?:png|jpg|jpeg|gif|png|svg))`)
	if err != nil {
		// panic(fmt.Errorf("error compiling regular expression: %w", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if !imageURLRegex.MatchString(film.ImageURL) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid image URL format",
		})
		return
	}

	err = repository.UpdateFilm(database.DbConnection, film)

	if err != nil {
		// panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Film",
	})
}

func DeleteFilm(c *gin.Context) {
	var film models.Film
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	film.IDFilm = int(id)

	err = repository.DeleteFilm(database.DbConnection, film)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Film",
	})
}

func GetFilmByFilmID(c *gin.Context) {
	var (
		result gin.H
		filmID int
	)

	filmIDParam := c.Param("id")
	filmID, err := strconv.Atoi(filmIDParam)
	if err != nil {
		result = gin.H{
			"error": "Invalid film ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	films, err := repository.GetFilmByID(database.DbConnection, filmID)
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
