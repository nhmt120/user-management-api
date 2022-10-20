package main

import (
	"UserManagementAPI/inits"
	"UserManagementAPI/models"
	"UserManagementAPI/utils"
)

// @title          Swagger User Management API
// @version        1.0
// @description    This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization

func main() {
	db := utils.ConnectDatabase()
	db.AutoMigrate(&models.User{})
	enforcer := inits.SetPolicies(db)
	inits.SetRoutes(db, enforcer)
}
