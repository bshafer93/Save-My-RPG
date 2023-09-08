package smrpg

// import the package we need to use
import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"savemyrpg/dal"
	"time"

	_ "github.com/lib/pq"
)

func RetrieveAllCampaignSaves(w http.ResponseWriter, r *http.Request) {

	resp_bytes, err := io.ReadAll(r.Body)
	fmt.Println(resp_bytes)
	fmt.Println(string(resp_bytes))
	if err != nil {
		fmt.Println(err)
	}
	campaignInfoJson := &CampaignID{}

	err = json.Unmarshal(resp_bytes, campaignInfoJson)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("GroupID: ", campaignInfoJson.Group_ID)

	saves := dal.GetAllCampaignSaves(campaignInfoJson.Group_ID)

	savesJson, err := json.Marshal(saves)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(savesJson)
}

func SaveImageUploadHandler(w http.ResponseWriter, r *http.Request) {
	resp_bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	file_name := r.Header.Get("file_name")
	campaign_id := r.Header.Get("group_id")

	BunnyUploadFile(campaign_id, file_name, resp_bytes)

	w.Write([]byte("Save Image UPLOADED"))
}

func SaveUploadHandler(w http.ResponseWriter, r *http.Request) {
	resp_bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	save_owner := r.Header.Get("email")
	file_name := r.Header.Get("file_name")
	folder_name := r.Header.Get("save_folder_name")
	campaign_id := r.Header.Get("group_id")

	BunnyUploadFile(campaign_id, folder_name+"/"+file_name, resp_bytes)

	save := dal.Save{}
	h := sha256.New()
	h.Write(resp_bytes)
	url := "https://ny.storage.bunnycdn.com/savemyrpg/bg3_saves/" + campaign_id + "/" + folder_name + "/" + file_name
	save.Folder_Name = folder_name
	save.Hash = fmt.Sprintf("%x", h.Sum(nil))
	save.Group_ID = campaign_id
	save.Save_Owner = save_owner
	save.CDN_Path = url
	save.Date_Created = time.Now().Format("2 Jan 2006 15:04:05")

	dal.AddSave(&save)

	w.Write([]byte("fILE UPLOADED"))

}
