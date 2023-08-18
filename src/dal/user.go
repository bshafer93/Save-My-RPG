package dal

import (
	"fmt"
)

type User struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

func AddUser(user *User) bool {
	_, err := db.Exec(`INSERT INTO users ("username","email") VALUES($1, $2)`, user.Username, user.Email)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to add user")
		return false
	}
	return true
}

func RemoveUser(email string) bool {
	q := `DELETE FROM users WHERE email =$1`
	_, err := db.Exec(q, email)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to remove user")
		return false
	}

	return true
}

func UpdateUserEmail() {

}

func UpdateUserName() {

}

func FindUserEmail(email string) bool {
	var doesExist bool

	//q := `SELECT EXISTS (SELECT FROM users WHERE email='$1');`
	qu := fmt.Sprintf(`SELECT EXISTS (SELECT FROM users WHERE email='%s');`, email)
	fmt.Println("Query: ", qu)
	db.QueryRow(qu).Scan(&doesExist)

	fmt.Println("FindUser Result: ", doesExist)
	return doesExist
}

func GetUser(email string) *User {
	var u User
	q := "SELECT email,username FROM users WHERE email = '" + email + "';"

	row := db.QueryRow(q)
	fmt.Print()
	row.Scan(&u.Email, &u.Username)
	fmt.Print(u.Email)
	return &u
}
