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
	inits.SetRoutes(db, &enforcer)
}

// var users = []models.User{
// 	{ID: 1, Name: "Vince", Email: "vince@gmail.com", Password: "1234", Role: "Admin", Status: "Active", CreatedAt: time.Now(), UpdatedAt: time.Now()},
// 	{ID: 2, Name: "Marc", Email: "marc@gmail.com", Password: "4567", Role: "User", Status: "Active", CreatedAt: time.Now(), UpdatedAt: time.Now()},
// }
