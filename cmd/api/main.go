package main

import (
	"github.com/gin-gonic/gin"

	"github.com/MitSonani/go-url-shortner/internal/db"
	"github.com/MitSonani/go-url-shortner/internal/handler"
)

func main() {

	db.ConnectDB()

	r := gin.Default()

	r.GET("/health", handler.HealthHandler)
	r.POST("/shorten", handler.ShortenURL)
	r.GET("/:code", handler.RedirectURL)

	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
