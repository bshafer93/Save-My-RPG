package smrpg

// import the package we need to use
import (
	"archive/zip"
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

func SaveUploadHandler(w http.ResponseWriter, r *http.Request) {
	//500KB Chunks
	buf := make([]byte, 512000)

	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			// Process or save the chunk data here...
			fmt.Printf("Received %d bytes\n", n)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Failed reading request body", http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte("Upload received!"))
}

func Init() bool {

	_, err := LoadConfiguration("/go/src/savemyrpgserver/config.json")
	if err != nil {
		return false
	}

	cert, err := tls.LoadX509KeyPair(config.SERVER_CERT, config.SERVER_KEY)
	if err != nil {
		return false
	}

	tls_config := &tls.Config{Certificates: []tls.Certificate{cert}}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/serverinfo", ServerInfoHandler)
	mux.HandleFunc("/getfullsave", SendFullFile)
	mux.HandleFunc("/login", Login)

	server = &http.Server{
		Addr:              ":" + config.SERVER_PORT,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		TLSConfig:         tls_config,
		Handler:           mux,
	}

	if !dal.Init() {
		return false
	}

	return true
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

func Register(username string, email string) (*User, error) {
	// Check if user already exists
	if dal.FindUserEmail(email) {
		return nil, errors.New("user email taken")
	}
	new_user := User{Username: username, Email: email}
	// Hash the password

	// Save user to database
	if !dal.AddUser(&new_user) {
		return nil, errors.New("user Could not be added")
	}
	// Return the user or error
	return &new_user, nil
}

func Login(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println("Username: "+userInfoJson.Username+"\nEmail:", userInfoJson.Email)

	// Check if user exists
	if !dal.FindUserEmail(userInfoJson.Email) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username does not exist"))
	}

	userInfo := dal.GetUser(userInfoJson.Email)
	tokenString := CreateLoginToken(userInfo)

	println("User: " + userInfo.Username + " Logged in!")

	w.Header().Add("jwt-token", tokenString)
	w.Write([]byte("Logged in!"))
}

func ZipSaveFile(folder_name string) string {

	save_file_path := "/go/src/savemyrpgserver" + config.SAVES_PATH + folder_name
	zip_file_path := save_file_path + ".zip"

	file_list, err := os.ReadDir(save_file_path)
	PrintError(err)

	fmt.Println("creating zip archive...")

	archive, err := os.Create(zip_file_path)
	PrintError(err)

	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	for _, f := range file_list {
		file_name := f.Name()
		full_file_path := save_file_path + "/" + file_name

		fmt.Println("opening file..." + file_name)
		f1, err := os.Open(full_file_path)
		PrintError(err)
		defer f1.Close()

		fmt.Println("writing file to archive...")
		w1, err := zipWriter.Create(file_name)
		PrintError(err)

		if _, err := io.Copy(w1, f1); err != nil {
			PrintError(err)
		}

	}

	fmt.Println("Finished Zipping : " + folder_name + ".zip")
	zipWriter.Close()

	return zip_file_path

}

func GetZipFileBytes(full_path string) []byte {

	bytes, err := os.ReadFile(full_path)
	PrintError(err)
	return bytes
}

func SendZipFile(save_zip_path string) {
	const BufferSize = 512000
	save_file_path := "/go/src/savemyrpgserver" + config.SAVES_PATH + save_zip_path

	zip_file, err := os.Open(save_file_path)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer zip_file.Close()
	zip_file_info, err := zip_file.Stat()
	PrintError(err)
	zip_file_size := zip_file_info.Size()
	total_chunks := zip_file_size / BufferSize

	current_chunk := 1
	for {
		sfc := SaveFileChunk{}
		sfc.Data = make([]byte, BufferSize)
		bytesread, err := zip_file.Read(sfc.Data)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			fmt.Println("All chunks processed")
			break
		}

		sfc.ChunkSize = uint32(bytesread)
		sfc.TotalChunks = uint32(total_chunks)
		sfc.ChunkNumber = uint32(current_chunk)
	}

}

func JoinCampaign(w http.ResponseWriter, r *http.Request) {

}

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
