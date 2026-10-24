package main

import (
	"golang-dasar/config"
	"golang-dasar/handler"
	"golang-dasar/user"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	// {
	// 	router.GET("/", func(c *gin.Context) {
	// 		var users []user.User
	// 		config.DB.Find(&users)

	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": "berhasil",
	// 			"data":    users,
	// 		})
	// 	})
	// }

	// userRepository := user.NewRepository(config.DB)
	// user := user.User{
	// 	Name: "hilman",

	// }
	// userRepository.Save(user)

	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		//user auth
		api.POST("/register", userHandler.RegisterUser)
		api.POST("/login", userHandler.LoginUser)

	}

	router.Run()

	// layer structure
	/**
	- input  => tipe data utk inputan
	- handler mapping input ke struct => controller balikan response
	- service mapping ke struct User =>  bisnis logic
	- repository save struct user ke db  => model (database)

	*/

}
