package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env only if it exists (ignore missing file)
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	dbURL := os.Getenv("DB_URL")

	fmt.Println("Server running on port:", port)
	fmt.Println("Database URL:", dbURL)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":    port,
			"db_url":  dbURL,
			"message": "Hello oscar from your Go server!",
		})
	})

	r.Run(":" + port)
}
