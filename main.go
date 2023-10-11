package main

import (
	"log"
	"startup-anas/auth"
	"startup-anas/handler"
	"startup-anas/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Import the MySQL driver

func main() {
	
	dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checker", userHandler.CheckEmailAvailablity)
	api.POST("/avatars", userHandler.UploudAvatar)
	router.Run()
	
}
