package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/shortURL/internal/storage"
)

func GetOriginalURL(c *gin.Context, shortURL string, db *sql.DB) {
	originalURL, err := storage.GetOriginalURL(shortURL, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.Redirect(http.StatusFound, originalURL)
}
