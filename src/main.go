// main.go

package main

// import the package we need to use
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"savemyrpg/smrpgdal"
	"time"

	_ "github.com/lib/pq"
)

type ServerInfo struct {
	Name     string    `json:"Name"`
	LoggedAt time.Time `json:"LoggedAt"`
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

func main() {

	smrpgdal.DB_Connect()

	http.HandleFunc("/", handler)
	http.HandleFunc("/serverinfo", jsonHandler)
	log.Fatal(http.ListenAndServe(":8100", nil))
}
