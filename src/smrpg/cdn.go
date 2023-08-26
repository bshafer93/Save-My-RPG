package smrpg

import (
	"bytes"
	"context"
	"crypto"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
)

var bunny_config *bunnystorage.Config

func InitBunny() {

	readOnlyKey := config.BUNNYNET_READ_API_KEY
	readWriteKey := config.BUNNYNET_WRITE_API_KEY

	// Create new Config to be initialize a Client.
	bunny_config = &bunnystorage.Config{
		Application: &bunnystorage.Application{
			Name:    "Save My Rpg",
			Version: "0.0.1",
			Contact: "bshafer93@gmail.com",
		},
		StorageZone: "savemyrpg",
		Key:         readWriteKey,
		ReadOnlyKey: readOnlyKey,
		Endpoint:    bunnystorage.EndpointNewYork,
	}

}

func BunnyDownloadFile(save_name string) {
	// Create a new Client instance with the given Config.
	bunnyclient, err := bunnystorage.NewClient(bunny_config)
	if err != nil {
		fmt.Println(err)
	}

	buf, resp, err := bunnyclient.Download(context.Background(), "/TempFolder", "Norbertle-3812316722__TheHollow.zip")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Header)
	os.WriteFile("Norbertle-3812316722__TheHollow.zip", buf, 0777)
}

func BunnyUploadFile(dest_folder string, file_name string, file []byte) {

	url := "https://ny.storage.bunnycdn.com/savemyrpg/bg3_saves/" + dest_folder + "/" + file_name

	payload := bytes.NewReader(file)
	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("AccessKey", config.BUNNYNET_WRITE_API_KEY)
	req.Header.Add("content-type", "application/octet-stream")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

}

func BunnyUploadFileURL(dest_folder string, file_name string, file []byte) {

	url := "https://ny.storage.bunnycdn.com/savemyrpg/bg3_saves" + dest_folder + "/" + file_name

	payload := bytes.NewReader(file)
	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("AccessKey", config.BUNNYNET_WRITE_API_KEY)
	req.Header.Add("content-type", "application/octet-stream")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

}

func BunnyGenerateUploadFileURL(group_folder string, save_name string) string {

	path := "/savemyrpg/bg3_saves/" + save_name
	/*
		var securityKey = "token_security_key";
		var path = "/pathto/file.jpg";

		// Load the current time
		var unixBaseTime = new DateTime(1970,1,1,0,0,0,0,System.DateTimeKind.Utc);
		var currentTime = ((long)(DateTime.UtcNow - unixBaseTime).TotalSeconds);

		// Set the time of expiry to one hour from now
		var expires = currentTime + 3600;

		// Generate the token
		System.Security.Cryptography.MD5 md5 = System.Security.Cryptography.MD5.Create();

		string hashableBase = securityKey + path + expires;

		// If using IP validation
		// hashableBase += "146.14.19.7";

		byte[] outpufBuffer = md5.ComputeHash(Encoding.UTF8.GetBytes(hashableBase));
		var token = Convert.ToBase64String(outpufBuffer);
		token = token.Replace("\n", "").Replace("+", "-").Replace("/", "_").Replace("=", "");

		// Generate the URL
		var url = $"https://myzone.b-cdn.net{path}?token={token}&expires={expires}";
	*/

	unixExpireTime := time.Now().Unix() * 86400

	md5Hash := crypto.MD5.New()
	hashTable := fmt.Sprintf("%s%s%d", config.BUNNYNET_TOKEN_KEY, path, unixExpireTime)

	buffer := []byte(hashTable)
	token := b64.StdEncoding.EncodeToString(md5Hash.Sum(buffer))
	token = strings.Replace(token, "\n", "", -1)
	token = strings.Replace(token, "+", "-", -1)
	token = strings.Replace(token, "/", "_", -1)
	token = strings.Replace(token, "=", "", -1)

	url := fmt.Sprintf(`https://ny.storage.bunnycdn.com%s?token=%s&expires=%d`, path, token, unixExpireTime)
	return url
}

func BunnyListAllFIles() {
	// Create a new Client instance with the given Config.
	bunnyclient, err := bunnystorage.NewClient(bunny_config)
	if err != nil {
		fmt.Println(err)
	}
	// List files in the storage zone.
	files, _, err := bunnyclient.List(context.Background(), "/")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		log.Printf("File: %s, Size: %d\n", file.ObjectName, file.Length)
	}
}

func GetCheckSum(file_path string) hash.Hash {
	f, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h
}
