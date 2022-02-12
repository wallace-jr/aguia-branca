package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wallace-jr/aguia-branca/server/database"
	"github.com/wallace-jr/aguia-branca/server/models"
)

func ShowTrip(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400. gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var trip models.Trip

	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, trip)
}

