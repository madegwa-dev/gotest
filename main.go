package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	fmt.Println("Server running on port:", port)
	fmt.Println("Database URL:", dbURL)

	r := gin.Default()

	// Define root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":    port,
			"db_url":  dbURL,
			"message": "Hello oscar from your Go server!",
		})
	})

	// Start server
	r.Run(":" + port) // listens on the port from .env
}
