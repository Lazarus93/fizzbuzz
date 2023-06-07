package main

import (
	"fizzbuzz/config"
	_ "fizzbuzz/docs"
	"fizzbuzz/internal/fizzbuzz"
	"fizzbuzz/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

// @title FizzBuzz API
// @version 1.0
// @description This is a sample server for a FizzBuzz API.
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load config
	cfg := config.LoadConfig()

	r := gin.Default()

	fbService := fizzbuzz.NewService()
	router.SetupRoutes(r, fbService)

	url := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", cfg.Host, cfg.Port))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()
}
