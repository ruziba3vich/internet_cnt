package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ruziba3vich/shortURL/internal/handlers"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "shorturls")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbName := "urls"

	name := "../internal/db/" + dbName + ".sql"
	sqlFile, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlFile))
	log.Println(string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	r.POST("/create-short-url", func(c *gin.Context) {
		handlers.GenerateShortURL(c, db)
	})
	r.GET("/get/:shortURL", func(c *gin.Context) {
		shortURL := c.Param("shortURL")
		handlers.GetOriginalURL(c, shortURL, db)
	})

	address := "localhost:7777"
	log.Println("Server is listening on", address)
	if err := r.Run(address); err != nil {
		log.Fatal(err)
	}
}
