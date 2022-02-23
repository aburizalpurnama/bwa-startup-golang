package main

import (
	"bwastartup/user"
	"bwastartup/utils"
)

func main() {
	db, _ := utils.GetDb()

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{
		Name:       "Test",
		Email:      "test@mail.com",
		Occupation: "kang galer",
		Password:   "test",
	}

	userService.RegisterUser(userInput)
}
