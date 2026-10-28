package main

import (
	"golang-dasar/auth"
	"golang-dasar/config"
	"golang-dasar/handler"
	"golang-dasar/helper"
	"golang-dasar/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
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

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		//user auth
		api.POST("/register", userHandler.RegisterUser)
		api.POST("/login", userHandler.LoginUser)
		api.POST("/email_checkers", userHandler.CheckEmailAvailability)
		api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

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

// ambil nilai header Authorization Header -> Bearer token
// dr header authorization kita ambil tokennya saja
// validasi token -> ambil userId -> ambil user di DB berdasarkan userId lewat service
// set context isinya user
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//Bearer token
		tokenString := ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			tokenString = arrToken[1]
		}

		//validasi token
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//ubah  jwt.claims jadi int
		userID := int(claim["user_id"].(float64))

		//cari data user berdasarkan userId
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//set context isinya user -> currentUser (user yg sedang aktif saat ini)
		c.Set("currentUser", user)

	}
}
