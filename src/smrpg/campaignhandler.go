package smrpg

// import the package we need to use
import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"savemyrpg/dal"

	_ "github.com/lib/pq"
)

type CampaignID struct {
	Group_ID string `json:"id"`
}

func RetrieveAllJoinedCampaigns(w http.ResponseWriter, r *http.Request) {
	resp_bytes, err := io.ReadAll(r.Body)
	fmt.Println(resp_bytes)
	fmt.Println(string(resp_bytes))
	if err != nil {
		fmt.Println(err)
	}
	userInfoJson := &User{}

	err = json.Unmarshal(resp_bytes, userInfoJson)

	if err != nil {
		fmt.Println(err)
	}

	campaigns := dal.GetAllJoinedCampaigns(userInfoJson.Email)

	if campaigns == nil {
		fmt.Println("No Campaigns Found")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No Campaigns Found"))
	}

	campainsJson, err := json.Marshal(campaigns)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(campainsJson)

}

func RetrieveCampaign(w http.ResponseWriter, r *http.Request) {
	resp_bytes, err := io.ReadAll(r.Body)
	fmt.Println(resp_bytes)
	fmt.Println(string(resp_bytes))
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}

	err = json.Unmarshal([]byte(resp_bytes), &result)
	if err != nil {
		fmt.Println(err)
	}

	campaign := dal.GetCampaignInfo(result["id"].(string))

	if campaign == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad ID"))
		return
	}
	campaignJson, err := json.Marshal(campaign)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Server Error"))
		return
	}

	w.Write(campaignJson)

}

func GetAvailablePlayerSlot(group_id string) (string, error) {
	campaign := dal.GetCampaignInfo(group_id)

	if campaign.Player02_Email == nil || *campaign.Player02_Email == "" {
		return `player02_email`, nil
	}

	if campaign.Player03_Email == nil || *campaign.Player03_Email == "" {
		return `player03_email`, nil
	}

	if campaign.Player04_Email == nil || *campaign.Player04_Email == "" {
		return `player04_email`, nil
	}

	return "", errors.New("All Campaign Slots full...")

}

func UserJoinCampaign(w http.ResponseWriter, r *http.Request) {
	resp_bytes, err := io.ReadAll(r.Body)
	fmt.Println(resp_bytes)
	fmt.Println(string(resp_bytes))
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}

	err = json.Unmarshal([]byte(resp_bytes), &result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result["Email"], result["id"])
	group_id := result["id"].(string)
	email := result["Email"].(string)
	slot, err := GetAvailablePlayerSlot(result["id"].(string))

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(err.Error()))
		return
	}
	err = dal.PlayerJoinCampaign(group_id, email, slot)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Joined Campaign!"))

}

func UserCreateCampaign(w http.ResponseWriter, r *http.Request) {
	resp_bytes, err := io.ReadAll(r.Body)
	fmt.Println("Create Campaign" + string(resp_bytes))
	fmt.Println(string(resp_bytes))
	if err != nil {
		fmt.Println(err)
	}

	groupInfoJson := &dal.Campaign{}

	err = json.Unmarshal(resp_bytes, groupInfoJson)

	if err != nil {
		fmt.Println(err)
	}

	added := dal.AddCampaign(groupInfoJson.Name, groupInfoJson.Host_Email)

	if !added {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to create campaign"))
		return
	}

	campaign := dal.GetCampaignInfoHostAndName(groupInfoJson.Host_Email, groupInfoJson.Name)

	if campaign == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to retrieve campaign..."))
		return
	}

	campaignJson, _ := json.Marshal(campaign)

	w.Write(campaignJson)

}
