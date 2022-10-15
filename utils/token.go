package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"exp":        time.Now().Add(time.Hour * 3).Unix(),
		"authorized": true,
		"email":      email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t, err

}

func ValidateToken(token string) (*jwt.Token, error) {
	//2nd arg function return secret key after checking if the signing method is HMAC and returned key is used by 'Parse' to decode the token)
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
