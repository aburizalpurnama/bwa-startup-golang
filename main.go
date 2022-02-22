package main

import (
	"bwastartup/user"
	"bwastartup/utils"
)

func main() {
	db, _ := utils.GetDb()

	userRepository := user.NewRepository(db)
	user := user.User{
		Name: "Test",
	}

	userRepository.Save(user)
}
