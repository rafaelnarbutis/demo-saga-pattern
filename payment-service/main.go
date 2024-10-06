package main

import (
	"net/http"
	"payment-service/infra"
	"payment-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/payments", payments)

	infra.InitConfig()

	router.Run(":8080")
}

func payments(c *gin.Context) {
	var payment models.Payment

	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{})

	go infra.SendMessage(payment)
}
