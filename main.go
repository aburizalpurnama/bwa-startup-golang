package main

import (
	"bwastartup/user"
	"bwastartup/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users", UserHandler)

	router.Run()
}

func UserHandler(c *gin.Context) {
	db, _ := utils.GetDb()

	var users []user.User

	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
