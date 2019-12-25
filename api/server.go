package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Server struct {
	Router *gin.Engine
}

func (s *Server) Run(port string) error {
	server.Router = gin.Default()
	server.configureRoutes()

	return s.Router.Run(port)
}

const (
	kPort string = "TODO_APP"
)

var server = Server{}

func init() {
	// Load environment variables from .env files
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func Run() {
	port := os.Getenv(kPort)
	if port == "" {
		port = ":8080"
	}

	if err := server.Run(port); err != nil {
		log.Println("Failed to start server with error: ")
		log.Fatal(err)
	}
}
