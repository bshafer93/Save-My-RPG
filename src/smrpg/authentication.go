package smrpg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

func Login(w http.ResponseWriter, r *http.Request) {

	resp_bytes, err := io.ReadAll(r.Body)
	fmt.Println(resp_bytes)
	fmt.Println(string(resp_bytes))
	if err != nil {
		fmt.Println(err)
	}
	userInfoJson := &User{}

	err = json.Unmarshal(resp_bytes, userInfoJson)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Username: "+userInfoJson.Username+"\nEmail:", userInfoJson.Email)

	// Check if user exists
	if !dal.FindUserEmail(userInfoJson.Email) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username does not exist"))
	}

	userInfo := dal.GetUser(userInfoJson.Email)
	tokenString := CreateLoginToken(userInfo)

	println("User: " + userInfo.Username + " Logged in!")

	w.Header().Add("jwt-token", tokenString)
	w.Write([]byte("Logged in!"))
}

func Register(username string, email string) (*User, error) {
	// Check if user already exists
	if dal.FindUserEmail(email) {
		return nil, errors.New("user email taken")
	}
	new_user := User{Username: username, Email: email}
	// Hash the password

	// Save user to database
	if !dal.AddUser(&new_user) {
		return nil, errors.New("user Could not be added")
	}
	// Return the user or error
	return &new_user, nil
}
