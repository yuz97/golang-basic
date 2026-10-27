package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error) //string adalah token
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

// tidak memerlukan dependency
func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 1).Unix(), //expired 24 jam = 1 hari
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return tokenString, err
	}

	return tokenString, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {

	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(secret), nil

	})

	if err != nil {
		return token, err
	}

	return token, nil
}
