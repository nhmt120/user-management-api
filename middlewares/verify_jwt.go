package middlewares

import (
	"UserManagementAPI/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyJWT() gin.HandlerFunc {
	utils.WriteLog("AppLog.txt")

	return func(context *gin.Context) {
		const BearerSchema string = "Bearer"
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("No Authorization header found.")
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found."})
			return
		}

		tokenString := authHeader[len(BearerSchema):]

		b64data := tokenString[strings.IndexByte(tokenString, ' ')+1:]

		if token, err := utils.ValidateToken(b64data); err != nil {
			fmt.Println("Invalid token: ", tokenString)
			log.Println(err.Error())

			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token"})
			return
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				context.AbortWithStatus(http.StatusUnauthorized)
				return
			} else {
				if token.Valid {
					context.Set("email", claims["email"])
					return
				} else {
					context.AbortWithStatus(http.StatusUnauthorized)
					return
				}
			}
		}
	}
}
