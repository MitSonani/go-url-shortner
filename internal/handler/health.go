package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var data = []string{"hello", "world"}

func HealthHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Health is good",
		"data":    data,
	})

}
