package dal

import (
	"fmt"
	"strconv"
)

type Campaign struct {
	id         int    `json:"id"`
	name       string `json:"name"`
	host_email string `json:"host_email"`
	p2_email   string `json:"p2_email"`
	p3_email   string `json:"p3_email"`
	p4_email   string `json:"p4_email"`
}

func AddCampaign(campaign_name string, host_email string, player_emails [3]string) bool {
	_, err := db.Exec(`INSERT INTO groups ("name","host_email","player02_email","player03_email","player04_email") VALUES ($1,$2,$3,$4,$5)`, campaign_name, host_email, player_emails[0], player_emails[1], player_emails[2])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to add user")
		return false
	}
	return true
}

func RemoveCampaign(campaign_id int, host_email string) bool {
	q := `DELETE FROM groups WHERE id =$1 AND host_email= $2 `
	_, err := db.Exec(q, strconv.Itoa(campaign_id), host_email)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to remove campaign")
		return false
	}
	return true
}
