package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"bwastartup/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := utils.GetDb()

	if err != nil {
		log.Fatal(err.Error)
	}

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	// API versioning
	api := router.Group("/api/v1")

	api.POST("/user", userHandler.RegisterUser)

	router.Run()
}
