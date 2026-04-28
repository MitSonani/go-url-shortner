package handler

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/MitSonani/go-url-shortner/internal/model"
	"github.com/MitSonani/go-url-shortner/internal/store"
	"github.com/MitSonani/go-url-shortner/internal/utils"
	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {

	var req model.ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
	}

	parsedURL, err := url.ParseRequestURI(req.URL)

	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid URL",
		})
		return
	}

	shortCode := utils.GenerateShoreCode(6)

	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		req.URL = "http://" + req.URL
	}

	for code, url := range store.URLStore {
		if url == req.URL {
			c.JSON(http.StatusOK, gin.H{
				"short_url": code,
			})
			return
		}
	}

	store.Mu.Lock()
	store.URLStore[shortCode] = req.URL
	store.Mu.Unlock()

	baseURL := "http://localhost:8080"

	c.JSON(http.StatusCreated, gin.H{
		"message":   " URL Created",
		"short_url": baseURL + "/" + shortCode,
	})
}

func RedirectURL(c *gin.Context) {
	shortCode := c.Param("code")

	store.Mu.RLock()
	originalURL, exists := store.URLStore[shortCode]
	store.Mu.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "url not found",
		})
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}
