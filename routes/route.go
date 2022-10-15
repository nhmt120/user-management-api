package routes

import (
	"UserManagementAPI/controllers"
	"UserManagementAPI/middlewares"
	"UserManagementAPI/static"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(db *gorm.DB) {
	r := gin.Default()
	userController := controllers.NewUserController(db)
	authController := controllers.NewAuthController(db)

	r.POST("/register", userController.Register)
	r.POST("/login", authController.Login)

	authorized := r.Group("/")
	authorized.Use(middlewares.VerifyJWT())
	{
		authorized.POST("/update", userController.Update)
		authorized.GET("/get-all", userController.GetAll)
		authorized.GET("/get-by-email", userController.GetByEmail)
		authorized.DELETE("/:id", userController.Delete)
		authorized.DELETE("/delete-all", userController.DeleteAll)
	}

	// r.Run() // listen and serve on 0.0.0.0:8080
	r.Run(static.GIN_HOST + ":" + static.GIN_PORT)
}
