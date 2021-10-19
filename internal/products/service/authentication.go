package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Jti string `json:"jti"`
	Sub string `json:"sub"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])

		}

		var claims authCustomClaims
		_, err := jwt.ParseWithClaims(encodedToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("fThWmZq4t7w!z%C*F-JaNdRgUkXn2r5u1qe5tr"), nil
		})

		v, _ := err.(*jwt.ValidationError)

		if v.Errors == jwt.ValidationErrorExpired && claims.ExpiresAt > time.Now().Unix()-(86400*14) {
			log.Printf("Token expired at %d", claims.ExpiresAt)
			return nil, errors.New("token expired")
		}

		return []byte(service.secretKey), nil
	})

}
