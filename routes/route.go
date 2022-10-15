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
	// r.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{json.Marshal(controllers)
	// 	})
	// })

	controller := controllers.NewUserController(db)

	authorized := r.Group("/")

	authorized.Use(middlewares.VerifyJWT())
	{
		authorized.POST("/update", controller.Update)
	}

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.DELETE("/:id", controller.Delete)
	r.GET("/get-all", controller.GetAll)

	// r.Run() // listen and serve on 0.0.0.0:8080
	r.Run(static.GIN_HOST + ":" + static.GIN_PORT)
}
