package smrpg

import (
	"fmt"
	"savemyrpg/dal"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateLoginToken(user *dal.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user.Username
	claims["email"] = user.Email
	claims["logged-in"] = true
	claims["exp"] = time.Now().Add(24 * time.Hour)
	tokenString, err := token.SignedString([]byte(config.JWT_SECRET_KEY))
	if err != nil {
		fmt.Println("Failed token creation for: ", user.Email)
	}
	return tokenString
}

func VerifyJWT(token string) bool {

	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET_KEY), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println(err)
			return false
		}
	}

	if !tkn.Valid {
		fmt.Println("Token Not Valid")
		return false
	}
	return true
}
