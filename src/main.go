// main.go

package main

// import the package we need to use
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"savemyrpg/dal"
	"time"

	_ "github.com/lib/pq"
)

type ServerInfo struct {
	Name     string    `json:"Name"`
	LoggedAt time.Time `json:"LoggedAt"`
}
type Session struct {
	UserID    int64
	SessionID string
	// other fields like expiration time, etc.
}
type User = dal.User

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Save My RPG, my I take your order please?: %s", r.URL.Path[1:])
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	serverInfo := ServerInfo{}

	serverInfo.Name = "Home Server!"
	serverInfo.LoggedAt = time.Now()
	serverInfoJson, _ := json.Marshal(serverInfo)
	w.Write(serverInfoJson)
}

/*
	func printUser(u *User) {
		fmt.Printf("Name: %s \nEmail: %s \n------------\n", &u.username, &u.email)
	}
*/

func main() {
	Init()
	log.Fatal(Start())
}

func Init() bool {
	config = LoadConfiguration("config.json")
	b := dal.Init()
	if b == false {
		return false
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/serverinfo", jsonHandler)
	return true
}

func Start() error {
	err := http.ListenAndServe(":"+config.server_port, nil)
	return err
}

func Register(username string, email string) (*User, error) {
	// Check if user already exists
	if dal.FindUser(email) {
		return nil, errors.New("User Email Taken")
	}
	new_user := User{Username: username, Email: email}
	// Hash the password

	// Save user to database
	if !dal.AddUser(&new_user) {
		return nil, errors.New("User Could not be added")
	}
	// Return the user or error
	return &new_user, nil
}

func Login(username string, email string) (*Session, error) {
	// Check if user exists and retrieve their hashed password
	if !dal.FindUser(email) {
		return nil, errors.New("Username does not exist")
	}
	// Compare provided password with stored hashed password

	// If they match, create a new session

	// Return the session or error
	return nil, errors.New("function not implemented")
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
