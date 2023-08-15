package dal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() bool {
	return Connect()
}

func Connect() bool {
	connStr := "postgres://admin:ninjame@192.168.1.155:9432/default?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Connection Failed!")
		log.Fatal(err)
		return false
	}

	fmt.Println("Connection Established!")
	return true
}
