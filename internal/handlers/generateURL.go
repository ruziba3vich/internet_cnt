package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/shortURL/internal/models"
	"github.com/ruziba3vich/shortURL/internal/storage"
)

func GenerateShortURL(c *gin.Context, db *sql.DB) {
	var url models.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println(url.Url)
	newURL, err := storage.GenerateShortURL(url.Url, db)
	if err != nil {
		fmt.Println("--------------------------->")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, newURL)
}
