package smrpg

// import the package we need to use
import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"savemyrpg/dal"
	"time"

	_ "github.com/lib/pq"
)

type ServerInfo struct {
	Name     string    `json:"Name"`
	LoggedAt time.Time `json:"LoggedAt"`
}

type User = dal.User

var server *http.Server

func handler(w http.ResponseWriter, r *http.Request) {
	page_text := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title></title>
		<style>
			/* Style to ensure the text is always centered */
			html, body {
				height: 100%;
				margin: 0;
				font-family: Arial, sans-serif;
				display: flex;
				align-items: center;
				justify-content: center;
			}
		</style>
	</head>
	<body>
		<!-- Text to be centered -->
		<div>
			Hi Adam!
		</div>
	</body>
	</html>`
	w.Write([]byte(page_text))
}

func SendFullFile(w http.ResponseWriter, r *http.Request) {
	fullFile := FullSaveFileJson{}
	fullFile.FolderName = "Norbertle-31112316728_smrpg.zip"
	fullFile.Data = b64.StdEncoding.EncodeToString(GetZipFileBytes("/go/src/savemyrpgserver" + config.SAVES_PATH + fullFile.FolderName))
	fullFile.FolderSize = int64(len(fullFile.Data))
	fullFileJson, _ := json.Marshal(fullFile)
	w.Write(fullFileJson)
}

func ServerInfoHandler(w http.ResponseWriter, r *http.Request) {
	serverInfo := ServerInfo{}

	serverInfo.Name = "Home Server!"
	serverInfo.LoggedAt = time.Now()
	serverInfoJson, _ := json.Marshal(serverInfo)
	w.Write(serverInfoJson)
}

func Start() error {

	err := server.ListenAndServeTLS(config.SERVER_CERT, config.SERVER_KEY)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to start server...")
		return err
	}
	fmt.Println("TLS Connection Established!")
	return nil
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

}

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
