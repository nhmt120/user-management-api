package main

import (
	"UserManagementAPI/models"
	"UserManagementAPI/routes"
	"UserManagementAPI/utils"
)

// var users = []models.User{
// 	{ID: 1, Name: "Vince", Email: "vince@gmail.com", Password: "1234", Role: "Admin", Status: "Active", CreatedAt: time.Now(), UpdatedAt: time.Now()},
// 	{ID: 2, Name: "Marc", Email: "marc@gmail.com", Password: "4567", Role: "User", Status: "Active", CreatedAt: time.Now(), UpdatedAt: time.Now()},
// }

func main() {
	db := utils.ConnectDatabase()

	// router := gin.Default()
	// router.GET("/getAllUsers", getUsers)
	// router.Run("localhost:8080")

	db.AutoMigrate(&models.User{})

	routes.SetRoutes(db)
	// userRepository := repositories.NewUserRepository(db)

	// user := models.User{ID: 1, Name: "Vincent", Email: "vincent@gmail.com", Password: "1234"}
	// user := models.User{Name: "Vince", Email: "vince@gmail.com", Password: "1234"}

	// userRepository.UpdateUser(user)

	// userRepository.DeleteUser(1)

	// userRepository.CreateUser(user)

	// user := userRepository.GetUserById(2)
	// userJSON, err := json.Marshal(user)
	// fmt.Println(string(userJSON))
}
