package main

import (
	"log"

	"github.com/IkhsanDS/golang-api/database"
	_ "github.com/IkhsanDS/golang-api/docs" // Importing the generated docs package
	"github.com/IkhsanDS/golang-api/models"
	"github.com/IkhsanDS/golang-api/router"
	"github.com/gin-gonic/gin"
)

// @title Gin Modular API
// @version 1.0
// @description Contoh struktur modular dengan Gin + Swagger
// @BasePath /
// @schemes http
func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	// auto migrate
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	r := gin.Default()
	router.Setup(r, db)

	log.Println("http://localhost:8080/swagger/index.html")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
