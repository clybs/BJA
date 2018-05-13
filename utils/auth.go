package utils

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Auth struct {
	JwtMiddleware *jwtmiddleware.JWTMiddleware
	SecretKey     string "cliburn"
	SigningMethod jwt.SigningMethod
	Token         string
}

func (at *Auth) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (at *Auth) CreateToken(id string) string {
	// Embed Admin information to `token`
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["userid"] = id

	// Expire in 10 mins
	token.Claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	tokenString, err := token.SignedString([]byte(at.SecretKey))
	if err != nil {
		return "Error occured"
	}
	return tokenString
}

func (at *Auth) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func (at *Auth) Init() {
	at.JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(at.SecretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
