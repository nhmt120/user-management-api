package controllers

import (
	"UserManagementAPI/models"
	"UserManagementAPI/repositories"
	"UserManagementAPI/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// utils.WriteLog("Logs")
// file, .repo. := os.OpenFile("Logs/db-logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
// log.SetOutput(file)

type UserController interface {
	Register(*gin.Context)
	Login(context *gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetAll(*gin.Context)
}

type userController struct {
	repo repositories.UserRepository
}

func NewUserController(db *gorm.DB) UserController {
	repo := repositories.NewUserRepository(db)
	utils.WriteLog("AppLog.txt")
	return userController{repo: repo}
}

func (c userController) Register(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err == nil {
		hash_password, err_1 := utils.HashPassword(user.Password)
		if err_1 != nil {
			log.Println("Error: Register, hashing failed.")
		} else {
			user.Password = hash_password
		}
		err_0 := c.repo.CreateUser(user)
		if err_0 == nil {
			log.Println("Action success: Register User.")
			context.JSON(200, gin.H{
				"code":    1,
				"message": "Register user successfully.",
			})
		} else {
			log.Println(err_0.Error())
			log.Println("Action failed: Register User.")
			context.JSON(200, gin.H{
				"code":    1,
				"message": "Register user failed.",
			})
		}
	} else {
		log.Println("Gin failed: ", err.Error(), ".")
	}
}

func (c userController) Login(context *gin.Context) {
	data, _ := context.GetRawData()
	m := map[string]string{}
	json.Unmarshal(data, &m)

	// fmt.Println(string(data))
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

func (c userController) GetAll(context *gin.Context) {
	users, err := c.repo.GetAllUsers()

	if err == nil {
		log.Println("Action success: Get all user.")
		context.JSON(200, gin.H{
			"code":    1,
			"message": users,
		})
	} else {
		log.Println(err.Error())
		log.Println("Action failed: Get all user.")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "Get all user failed.",
		})
	}
}

func (c userController) Update(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err == nil {
		err_0 := c.repo.UpdateUser(user)
		if err_0 == nil {
			log.Println("Action success: Update user.")
			context.JSON(200, gin.H{
				"code":    1,
				"message": "Update user successfully.",
			})
			return
		} else {
			log.Println(err_0.Error())
			log.Println("Action failed: Update user.")
			context.JSON(200, gin.H{
				"code":    1,
				"message": "Update user failed.",
			})
		}
	} else {
		log.Println("Gin failed: ", err.Error(), ".")
	}
}

func (c userController) Delete(context *gin.Context) {
	id := context.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		context.JSON(200, gin.H{
			"code":    1,
			"message": "Gin failed: Delete user.",
		})
		return
	}

	_, err_1 := c.repo.GetUser(intID)
	if err_1 != nil {
		log.Println("Get user failed.")
		log.Println(err_1.Error())
		context.JSON(200, gin.H{
			"code":    0,
			"message": "User with ID = " + id + " not found.",
		})
		return
	}

	err_0 := c.repo.DeleteUser(intID)
	if err_0 == nil {
		log.Println("Action success: Delete user.")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "Delete user successfully.",
		})
	} else {
		log.Println(err_0.Error())
		log.Println("Action failed: Delete user.")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "Delete user failed.",
		})
	}
}
