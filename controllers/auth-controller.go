package controllers

import (
	"UserManagementAPI/repositories"
	"UserManagementAPI/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController interface {
	Login(context *gin.Context)
}

type authController struct {
	repo repositories.UserRepository
}

func NewAuthController(db *gorm.DB) AuthController {
	repo := repositories.NewUserRepository(db)
	utils.WriteLog("AppLog.txt")
	return authController{repo: repo}
}

func (c authController) Login(context *gin.Context) {
	data, _ := context.GetRawData()
	m := map[string]string{}
	json.Unmarshal(data, &m)

	email := m["email"]
	password := m["password"]
	if email == "" || password == "" {
		log.Println("Action failed: Login, missing email or password information.")
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "Missing login information.",
		})
		return
	}

	user, err := c.repo.GetUserByEmail(email)

	if err != nil {
		log.Println(err)
		log.Println("Action failed: Login, invalid email.")
		context.JSON(200, gin.H{
			"code":    -1,
			"message": "User with email " + email + " does not exist.",
		})
		return
	}
	if is_valid := utils.ComparePassword(user.Password, password); is_valid {
		log.Println("Action success: Login.")
		jwt_token, err := utils.GenerateJWT(user.Email)

		if err != nil {
			log.Println(err.Error())
			log.Println("Utils failed: Generate JWT token.")
		}
		context.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "Login successfully.",
			"token":   jwt_token,
		})
		return
	} else {
		log.Println("Action failed: Login, invalid password")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "Login failed: Invalid credentials."})
		return
	}
}
