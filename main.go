package main

import (
	"log"

	"github.com/IkhsanDS/golang-api/database"
	_ "github.com/IkhsanDS/golang-api/docs"
	"github.com/IkhsanDS/golang-api/models"
	"github.com/IkhsanDS/golang-api/router"
)

// @title           Todo API
// @version         1.1
// @description     Todo API with JWT Auth, RBAC, and pagination.
// @BasePath        /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description     Format: Bearer {token}
func main() {
	// Connect to the database

	database.Connect()
	if err := database.DB.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		log.Fatalf("migration error: %v", err)
	}
	r := router.Setup()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
