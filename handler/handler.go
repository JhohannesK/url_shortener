package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhohannesK/go-url-shortener/shortener"
	"github.com/jhohannesK/go-url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}


func CreateShortUrl(c *gin.Context){
	var creationRequest UrlCreationRequest

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:8080/"
	c.JSON(200, gin.H{
		"message": "short url created successfully",
		"short_url": host +shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context){
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}