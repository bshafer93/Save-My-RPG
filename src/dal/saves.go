package dal

import (
	"fmt"
)

type Save struct {
	Hash         string `json:"hash" db:"hash"`
	Group_ID     string `json:"group_id" db:"group_id"`
	Save_Owner   string `json:"save_owner" db:"save_owner"`
	Folder_Name  string `json:"folder_name" db:"folder_name"`
	Full_Path    string `json:"-" db:"full_local_path"`
	CDN_Path     string `json:"cdn_path" db:"cdn_path"`
	Date_Created string `json:"date_created" db:"date_created"`
}

func GetAllCampaignSaves(id string) []Save {
	q := fmt.Sprintf(`SELECT * FROM saves WHERE group_id = '%[1]s'`, id)
	saves := make([]Save, 0, 10)
	rows, err := db.Query(q)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		s := Save{}
		err = rows.Scan(&s.Hash,
			&s.Group_ID,
			&s.Save_Owner,
			&s.Folder_Name,
			&s.Full_Path,
			&s.CDN_Path,
			&s.Date_Created)
		if err != nil {
			fmt.Println(err)
		}
		saves = append(saves, s)
	}
	return saves
}

func AddSave(save *Save) bool {
	_, err := db.Exec(`INSERT INTO saves ("hash","group_id","save_owner","folder_name","cdn_path","date_created","full_local_path") VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		save.Hash,
		save.Group_ID,
		save.Save_Owner,
		save.Folder_Name,
		save.CDN_Path,
		save.Date_Created,
		"/")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to add save")
		return false
	}

	fmt.Println("New Save Added to DB")
	return true
}
