package dal

import (
	"errors"
	"fmt"
	"strconv"
)

type Campaign struct {
	ID             string  `json:"id" db:"id"`
	Name           string  `json:"name" db:"name"`
	Host_Email     string  `json:"host_email" db:"host_email"`
	Player02_Email *string `json:"p2_email" db:"player02_email"`
	Player03_Email *string `json:"p3_email" db:"player03_email"`
	Player04_Email *string `json:"p4_email" db:"player04_email"`
	Last_Save      *string `json:"last_save" db:"last_save"`
}

func UpdateCampaign(campaign_name string, host_email string, player_emails [3]string) bool {
	_, err := db.Exec(`INSERT INTO groups ("name","host_email","player02_email","player03_email","player04_email") VALUES ($1,$2,$3,$4,$5)`, campaign_name, host_email, player_emails[0], player_emails[1], player_emails[2])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to add user")
		return false
	}
	return true
}

func AddCampaign(campaign_name string, host_email string) bool {
	_, err := db.Exec(`INSERT INTO groups ("name","host_email") VALUES ($1,$2)`, campaign_name, host_email)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to add campaign")
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

func GetAllHostCampaigns(host_email string) []Campaign {

	q := fmt.Sprintf(`SELECT * FROM groups WHERE host_email = '%s'`, host_email)

	campaigns := make([]Campaign, 0, 10)
	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		c := Campaign{}
		err = rows.Scan(&c.ID,
			&c.Name,
			&c.Host_Email,
			&c.Player02_Email,
			&c.Player03_Email,
			&c.Player04_Email,
			&c.Last_Save)

		if err != nil {
			fmt.Println(err)
		}
		campaigns = append(campaigns, c)
	}

	return campaigns
}

func GetAllJoinedCampaigns(email string) []Campaign {
	q := fmt.Sprintf(`SELECT * FROM groups WHERE host_email = '%[1]s' OR player02_email = '%[1]s' OR player03_email = '%[1]s' OR player04_email = '%[1]s'`, email)
	campaigns := make([]Campaign, 0, 10)
	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		c := Campaign{}
		err = rows.Scan(&c.ID,
			&c.Name,
			&c.Host_Email,
			&c.Player02_Email,
			&c.Player03_Email,
			&c.Player04_Email,
			&c.Last_Save)
		fmt.Println("P4: ", c.Player04_Email)
		if err != nil {
			fmt.Println(err)
		}
		campaigns = append(campaigns, c)
	}

	return campaigns
}

func GetCampaignInfo(group_id string) *Campaign {
	q := fmt.Sprintf(`SELECT * FROM groups WHERE id = '%s'`, group_id)
	row := db.QueryRow(q)

	c := Campaign{}
	err := row.Scan(&c.ID,
		&c.Name,
		&c.Host_Email,
		&c.Player02_Email,
		&c.Player03_Email,
		&c.Player04_Email,
		&c.Last_Save)
	if err != nil {
		fmt.Println(err)
	}

	return &c

}

func GetCampaignInfoHostAndName(host_email string, name string) *Campaign {
	q := fmt.Sprintf(`SELECT * FROM groups WHERE host_email = '%s' AND name = '%s'`, host_email, name)

	row := db.QueryRow(q)

	c := Campaign{}
	err := row.Scan(&c.ID,
		&c.Name,
		&c.Host_Email,
		&c.Player02_Email,
		&c.Player03_Email,
		&c.Player04_Email,
		&c.Last_Save)
	if err != nil {
		fmt.Println(err)
	}

	return &c

}

func PlayerJoinCampaign(group_id string, email string, slot string) error {
	q := fmt.Sprintf(`UPDATE groups SET %s = '%s' WHERE id = '%s';`, slot, email, group_id)
	_, err := db.Exec(q)
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to add player to campaign")
	}

	return nil
}
