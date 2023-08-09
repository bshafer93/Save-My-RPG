package smrpgdal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DBConnect() bool {

	connStr := "postgres://admin:ninjame@192.168.1.33/default?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Connection Failed!")
		log.Fatal(err)
		return false
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
		fmt.Printf("Name: %s \nEmail: %s \n------------\n", u.username, u.email)
	}

	fmt.Println("Connection Established!")
	return true
}
