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
	connStr := "postgres://admin:ninjame@192.168.1.33/default?sslmode=disable"
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

func FindAll(q string) {
	/*
		rows, err := db.Query("SELECT * FROM users;")

		if err != nil {
			fmt.Println("Query Failed!")
			log.Fatal(err)
		}

		/*

			for rows.Next() {
				u := User{}
				err := rows.Scan(&u.email, &u.username)
				if err != nil {
					fmt.Println("Query Failed!")
				}
				fmt.Printf("Name: %s \nEmail: %s \n------------\n", u.username, u.email)
			}
	*/
	fmt.Printf(db.Stats().WaitDuration.Abs().String())
}
