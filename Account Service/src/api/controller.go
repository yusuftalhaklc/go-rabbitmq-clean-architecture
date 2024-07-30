package main

import (
	"account"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var body account.Register

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	err = usecase.Register(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server Error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
	})
}

func Verify(ctx *gin.Context) {
	code := ctx.Query("code")
	email := ctx.Query("email")

	if code == "" || email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Code and email are required"})
		return
	}
	err := usecase.Verify(account.VerifyCode{
		Email: email,
		Code:  code,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Account cannot verified",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Account verified successfully",
	})

}
