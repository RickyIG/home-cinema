package controllers

import (
	"home-cinema/database"
	"home-cinema/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserTicketHistory(c *gin.Context) {
	userID, isExist := c.Get("userID")
	if !isExist {
		c.JSON(http.StatusBadRequest, "User doesn't exist.")
		return
	}

	userTicketHistory, err := repository.GetUserTicketHistory(database.DbConnection, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := gin.H{
		"result": userTicketHistory,
	}
	c.JSON(http.StatusOK, result)
}

func GetUserTicketHistoryByTransactionID(c *gin.Context) {
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

	userTicketHistoryByID, err := repository.GetUserTicketHistoryByTransactionID(database.DbConnection, userID.(int), transaksiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := gin.H{
		"result": userTicketHistoryByID,
	}
	c.JSON(http.StatusOK, result)
}

func GetUserTicketHistoryByIDs(c *gin.Context) {
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

	ticketIDParam := c.Param("id_ticket")
	ticketID, err := strconv.Atoi(ticketIDParam)
	if err != nil {
		result := gin.H{
			"error": "Invalid transaksi ID",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	userTicketHistoryByID, err := repository.GetUserTicketHistoryByID(database.DbConnection, userID.(int), transaksiID, ticketID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := gin.H{
		"result": userTicketHistoryByID,
	}
	c.JSON(http.StatusOK, result)
}
