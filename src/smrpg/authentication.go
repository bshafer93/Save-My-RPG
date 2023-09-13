package smrpg

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"savemyrpg/dal"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

func CreateLoginToken(email string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["logged-in"] = true
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	tokenString, err := token.SignedString([]byte(config.JWT_SECRET_KEY))
	if err != nil {
		fmt.Println("Failed token creation for: ", email)
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

func AuthenticateJWTWrapper(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token, err := jwt.Parse(r.Header.Get("jwt-token"), func(token *jwt.Token) (interface{}, error) {

			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("You're Unauthorized!"))
				if err != nil {
					return nil, err

				}
			}

			return []byte(config.JWT_SECRET_KEY), nil
		})

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			_, err2 := w.Write([]byte("You're Unauthorized due to error parsing the JWT"))
			if err2 != nil {
				return
			}
		}

		if token.Valid {
			endpointHandler(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("You're Unauthorized due to invalid token"))
			if err != nil {
				return
			}
		}
	})

}

func AuthenticateJWT(tokenString string) (jwt.MapClaims, bool) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("not authorized")

		}
		return []byte(config.JWT_SECRET_KEY), nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}

	if !token.Valid {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}

	return nil, false
}

func Login(w http.ResponseWriter, r *http.Request) {

	_, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	user_email := r.Header.Get("email")
	pwd := r.Header.Get("pwd")

	if len(r.Header.Get("jwt-token")) > 0 {
		claims, ok := AuthenticateJWT(r.Header.Get("jwt-token"))

		if ok && fmt.Sprint(claims["email"]) == user_email {
			w.Write([]byte("Logged in!"))
			return
		}
	}

	fmt.Println("Email: "+user_email+"\npwd:", pwd)

	// Check if user exists
	if !dal.FindUserEmail(user_email) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password or username not good"))
	}

	//Check Password
	hashed_pwd := dal.GetPassword(user_email)

	if !CheckAndComparePassword(*hashed_pwd, []byte(pwd)) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Password or username not good"))
	}

	tokenString := CreateLoginToken(user_email)

	println("User: " + user_email + " Logged in!")

	w.Header().Add("jwt-token", tokenString)
	w.Write([]byte("Logged in!"))
}

func Register(username string, email string) (*User, error) {
	// Check if user already exists
	if dal.FindUserEmail(email) {
		return nil, errors.New("user email taken")
	}
	new_user := User{Pwd: username, Email: email}
	// Hash the password

	// Save user to database
	if !dal.AddUser(&new_user) {
		return nil, errors.New("user Could not be added")
	}
	// Return the user or error
	return &new_user, nil
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	resp_bytes, err := io.ReadAll(r.Body)
	fmt.Println(resp_bytes)
	fmt.Println(string(resp_bytes))
	if err != nil {
		fmt.Println(err)
	}

	var u User

	u.Email = r.Header.Get("email")

	pwd := r.Header.Get("pwd")

	fmt.Println("Email: "+u.Email+"\npwd:", pwd)

	// Check if email is already being used
	if dal.FindUserEmail(u.Email) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email taken"))
	}

	//Generate Hashed Password
	u.Pwd = HashAndSalt([]byte(pwd))

	if !dal.AddUser(&u) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Failed to Register User... Server Error!"))
	}

	println("User: " + u.Email + " registerd!")

	w.Write([]byte("user Registered! Please Login"))

}

func HashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(hash)
}

func CheckAndComparePassword(pwd_hashed string, pwd []byte) bool {

	byteHash := []byte(pwd_hashed)

	err := bcrypt.CompareHashAndPassword(byteHash, pwd)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
