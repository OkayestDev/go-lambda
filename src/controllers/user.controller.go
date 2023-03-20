package controllers

import (
	"learninggo/src/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(context *gin.Context) {
	var user = structs.User{
		Name: "d00d",
	}

	context.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}
