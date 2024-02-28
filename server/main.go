package main

import (
	"net/http"
	"os"
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.User{}, &models.Task{})
	// Check if the owner has already been created or not
	config.CreateOwnerAccount(db)


	// Create router from gin, will automatically import
	router := gin.Default()
	router.GET("/",func(c *gin.Context) {
		c.JSON(http.StatusOK,"Welcome to Test API")
	})

	// by default 8080
	// check your ip address in cmd, then type ipconfig
	// copy & paste IPv4 Address
	// Make sure put in .env 
	ipv4Address := os.Getenv("IPv4_ADDRESS")
	router.Run(ipv4Address + ":8080")
}