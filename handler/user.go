package handler

import (
	"bwastartup/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input ke struct RegisterUserInput
	// passing struct diatas kedalam service

	input := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		log.Fatal(err.Error)
		c.JSON(http.StatusBadRequest, input)
	}

	user, registErr := h.userService.RegisterUser(input)

	if registErr != nil {
		log.Fatal(registErr.Error)
		c.JSON(http.StatusBadRequest, input)
	}

	c.JSON(http.StatusOK, user)
}
