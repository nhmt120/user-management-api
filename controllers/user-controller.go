package controllers

import (
	"UserManagementAPI/models"
	"UserManagementAPI/repositories"
	"UserManagementAPI/utils"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	Register(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	DeleteAll(*gin.Context)
	GetAll(*gin.Context)
	GetByEmail(*gin.Context)
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

func (c userController) GetByEmail(context *gin.Context) {
	// users, err := c.repo.GetUserByEmail()
	// context.Request.URL.Query()
	email := context.Query("email")

	if email == "" {
		log.Println("Gin failed or empty email input.")
		context.JSON(200, gin.H{
			"code":    0,
			"message": "No email input received.",
		})
		return
	}

	user, err := c.repo.GetUserByEmail(email)

	if err == nil {
		log.Println("Action success: Get user with email = " + email + ".")
		context.JSON(200, gin.H{
			"code":    1,
			"message": user,
		})
	} else if user.Name == "" {
		log.Println("User with email = " + email + " does not exist.")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "User with email = " + email + " does not exist.",
		})
	} else {
		log.Println(err.Error())
		log.Println("Action failed: Get user by email.")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "Get user by email failed.",
		})
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

func (c userController) DeleteAll(context *gin.Context) {

	err_0 := c.repo.DeleteAll()
	if err_0 == nil {
		log.Println("Action success: Delete all users.")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "Delete all users successfully.",
		})
	} else {
		log.Println(err_0.Error())
		log.Println("Action failed: Delete all users.")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "Delete all users failed.",
		})
	}
}
