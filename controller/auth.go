package controller

import (
	"gotest_superindo/jwt"
	"gotest_superindo/pkg"
	"gotest_superindo/structs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	db, err := pkg.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	var user structs.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := db.Table("tb_user").Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Sukses register"})
}

func Login(context *gin.Context) {
	db, err := pkg.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	var request structs.Login
	var user structs.User

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := db.Table("tb_user").Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := jwt.GenerateJWT(user.Email, user.Nama)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success", "token": tokenString})
}
