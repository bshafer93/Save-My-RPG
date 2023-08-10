// main.go

package main

// import the package we need to use
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"savemyrpg/dal"
	"time"

	_ "github.com/lib/pq"
)

type ServerInfo struct {
	Name     string    `json:"Name"`
	LoggedAt time.Time `json:"LoggedAt"`
}

type User = dal.User
type SessionID = uint16

var loggedInUsers [65535]*User
var freeSessionID [65535]SessionID
var sessionIDTracker uint16 = 0

func main() {
	Init()
	log.Fatal(Start())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title></title>
    <style>
        /* Style to ensure the text is always centered */
        html, body {
            height: 100%;
            margin: 0;
            font-family: Arial, sans-serif;
            display: flex;
            align-items: center;
            justify-content: center;
        }
    </style>
</head>
<body>
    <!-- Text to be centered -->
    <div>
        `+"%s"+`
    </div>
</body>
</html>`, "Hi Adam!")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	serverInfo := ServerInfo{}

	serverInfo.Name = "Home Server!"
	serverInfo.LoggedAt = time.Now()
	serverInfoJson, _ := json.Marshal(serverInfo)
	w.Write(serverInfoJson)
}

func Init() bool {
	freeSessionID := rand.Perm(65535)
	sessionIDTracker = uint16(freeSessionID[0])

	_, err := LoadConfiguration("/go/src/savemyrpgserver/config.json")
	if err != nil {
		return false
	}

	if !dal.Init() {
		return false
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/serverinfo", jsonHandler)
	return true
}

func Start() error {
	err := http.ListenAndServeTLS(":"+config.SERVER_PORT, config.SERVER_CERT, config.SERVER_KEY, nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to start server...")
		return err
	}
	return nil
}

func Register(username string, email string) (*User, error) {
	// Check if user already exists
	if dal.FindUser(email) {
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

func Login(username string, email string) error {
	// Check if user exists and retrieve their hashed password
	if !dal.FindUser(email) {
		return errors.New("username does not exist")
	}

	user := dal.GetUser(email)

	loggedInUsers[sessionIDTracker] = user
	sessionIDTracker++

	// Compare provided password with stored hashed password

	// If they match, create a new session

	// Return the session or error
	return nil
}

func Logout(sessionID string) error {
	// Invalidate the session using the provided sessionID
	return errors.New("function not implemented")
}

func ResetPassword(username, newPassword string) error {
	// Validate user existence
	// Hash the new password
	// Update password in the database
	return errors.New("function not implemented")
}

func CheckAuthentication(sessionID string) (bool, error) {
	// Check if the sessionID is valid and not expired
	return false, errors.New("function not implemented")
}
