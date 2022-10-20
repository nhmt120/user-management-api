package inits

import (
	"UserManagementAPI/controllers"
	"UserManagementAPI/docs"
	"UserManagementAPI/middlewares"
	"UserManagementAPI/static"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRoutes(db *gorm.DB, enforcer *casbin.Enforcer) {

	// router := gin.Default()
	// router.GET("/getAllUsers", getUsers)
	// router.Run("localhost:8080")

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	userController := controllers.NewUserController(db)
	authController := controllers.NewAuthController(db)

	r.POST("/register", userController.Register(enforcer))
	r.POST("/login", authController.Login)

	authorized := r.Group("/")
	authorized.Use(middlewares.VerifyJWT())
	{
		//, middlewares.VerifyAccess("report", "write", enforcer)
		authorized.GET("/get-all", userController.GetAll)
		authorized.GET("/get-by-email", userController.GetByEmail)
		authorized.POST("/update", userController.Update)
		authorized.DELETE("/:id", userController.Delete)
		authorized.DELETE("/delete-all", userController.DeleteAll)
	}

	// r.Run() // listen and serve on 0.0.0.0:8080
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(static.GIN_HOST + ":" + static.GIN_PORT)
}
