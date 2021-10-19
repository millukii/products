package authentication

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					c.AbortWithStatus(http.StatusUnauthorized)
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				type jwtCustomClaims struct {
					Jti string `json:"jti"`
					Sub string `json:"sub"`
					jwt.StandardClaims
				}

				var claims jwtCustomClaims
				_, err := jwt.ParseWithClaims(bearerToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("JWT_SECRET")), nil
				})

				v, _ := err.(*jwt.ValidationError)

				if v.Errors == jwt.ValidationErrorExpired && claims.ExpiresAt > time.Now().Unix()-(86400*14) {
					log.Printf("Token expired at %d", claims.ExpiresAt)

					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}
				log.Printf("Token error %s", err)

				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			if token.Valid {
				log.Print("Token valid")

				c.Next()
			} else {
				log.Print("Token not valid")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			log.Print("Token is required")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}
