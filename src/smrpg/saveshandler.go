package smrpg

// import the package we need to use
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"savemyrpg/dal"

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
