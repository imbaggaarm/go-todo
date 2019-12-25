package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	kPort string = "GO_TO_PORT"
)

func main() {
	r := gin.Default()
	api := r.Group("/api/v1")

	// Configure Index
	api.GET("/", HandleIndex())

	// Load environment variables from .env files
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	port := os.Getenv(kPort)
	if port == "" {
		port = ":8080"
	}

	if err := r.Run(port); err != nil {
		log.Println("Failed to start server with error: ")
		log.Fatal(err)
	}
}

func HandleIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
