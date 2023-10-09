package main

import (
	"log"
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

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "tes simpan darfi service"
	// userInput.Email = "contoh@gmail.com"
	// userInput.Occupation = "anakband"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)

	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()
	
}
