package main

import (
	"golang-dasar/auth"
	"golang-dasar/campaign"
	"golang-dasar/config"
	"golang-dasar/handler"
	"golang-dasar/middleware"
	"golang-dasar/user"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//pangil godotenv utk memuat semua file .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("gagal load file .env: ", err)
	}

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

	//uji coba layer
	/**
	checkEmail, err := userService.IsEmailAvailable(user.CheckEmailInput{
		Email: "rudi@gmail.com",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(checkEmail)

	*/

	//repository
	userRepository := user.NewRepository(config.DB)
	campaignRepository := campaign.NewRepository(config.DB)

	//service
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()
	// fmt.Println(authService.GenerateToken(100)) //manual generate token
	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTMyMDIxMDcsInN1YiI6Mn0.M_UK44fEI98O2ffZaKRcA3Tj8GBuHTEzgyJFkvccwOc")
	// if err != nil {
	// 	fmt.Println("token tidak sama")
	// }

	// if token.Valid {
	// 	fmt.Println("berhasil")
	// } else {
	// 	fmt.Println("gagal!")
	// }

	//handler
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	//ambil gambar static
	router.Static("/images", "./images") // "./images" -> lokasi folder
	api := router.Group("/api/v1")
	{
		//user auth
		api.POST("/register", userHandler.RegisterUser)
		api.POST("/login", userHandler.LoginUser)
		api.POST("/email_checkers", userHandler.CheckEmailAvailability)
		api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

		//campaigns
		api.GET("/campaigns", campaignHandler.GetCampaigns)
		api.POST("/campaigns", middleware.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
		api.GET("/campaigns/:id", campaignHandler.GetCampaignByID)

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
