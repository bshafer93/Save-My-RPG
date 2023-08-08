// main.go

package main

// import the package we need to use
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type ServerInfo struct {
	Name     string    `json:"Name"`
	LoggedAt time.Time `json:"LoggedAt"`
}

type User struct {
	username string
	email    string
}

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

func printUser(u *User) {
	fmt.Printf("Name: %s \nEmail: %s \n------------\n", u.username, u.email)
}

func db_connect() {

	connStr := "postgres://admin:ninjame@192.168.1.33/default?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Connection Failed!")
		log.Fatal(err)
		os.Exit(0)
	}

	rows, err := db.Query("SELECT email,username FROM users;")

	if err != nil {
		fmt.Println("Query Failed!")
		log.Fatal(err)
		os.Exit(0)
	}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.email, &u.username)
		if err != nil {
			fmt.Println("Query Failed!")
		}
		printUser(&u)
	}

	fmt.Println("Connection Established!")
}

func main() {

	db_connect()

	http.HandleFunc("/", handler)
	http.HandleFunc("/serverinfo", jsonHandler)
	log.Fatal(http.ListenAndServe(":8100", nil))
}
