package middleware

import (
	"golang-dasar/auth"
	"golang-dasar/helper"
	"golang-dasar/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ambil nilai header Authorization Header -> Bearer token
// dr header authorization kita ambil tokennya saja
// validasi token -> ambil userId -> ambil user di DB berdasarkan userId lewat service
// set context isinya user

func AuthMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
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
