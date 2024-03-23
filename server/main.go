package main

import (
	"net/http"
	"os"
	"server/config"
	"server/controllers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.User{}, &models.Task{})
	// Check if the owner has already been created or not
	config.CreateOwnerAccount(db)

	// Controller
	userController := controllers.UserController{DB: db}
	taskController := controllers.TaskController{DB: db}
	
	// Router
	// Create router from gin, will automatically import
	router := gin.Default()
	router.GET("/",func(c *gin.Context) {
		c.JSON(http.StatusOK,"Welcome to Test API")
	})

	// User Router
	router.POST("/users/login", userController.Login)
	router.POST("/users", userController.CreateAccount)
	router.DELETE("/users/:id", userController.Delete)
	router.GET("/users/Employee", userController.GetEmployee)

	// Task Router
	router.POST("/tasks", taskController.Create)
	router.DELETE("/tasks/:id", taskController.Delete)
	router.PATCH("/tasks/:id/submit", taskController.Submit)
	// router.PATCH("/tasks/:id/reject", taskController.Reject)
	// router.PATCH("/tasks/:id/fix", taskController.Fix)
	// router.PATCH("/tasks/:id/approve", taskController.Approve)
	// router.GET("/tasks/:id", taskController.FindById)
	// router.GET("/tasks/review/asc", taskController.NeedToBeReview)
	// router.GET("/tasks/progress/:userId", taskController.ProgressTasks)
	// router.GET("/tasks/stat/:userId", taskController.Statistic)
	// router.GET("/tasks/user/:userId/:status", taskController.FindByUserAndStatus)

	// by default 8080
	// check your ip address in cmd, then type ipconfig
	// copy & paste IPv4 Address
	// Make sure put in .env 
	ipv4Address := os.Getenv("IPv4_ADDRESS")
	router.Run(ipv4Address + ":8080")
}