package smrpg

import (
	"context"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

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

func BunnyUploadFile(dest_folder string, dest_name string, source_save_name string, source_full_path string) {

	url := "https://ny.storage.bunnycdn.com/savemyrpg/" + dest_folder + "/" + dest_name

	b64_file := b64.StdEncoding.EncodeToString(GetZipFileBytes(source_full_path))

	payload := strings.NewReader("data:application/x-zip-compressed;name=" + source_save_name + ";base64," + b64_file)

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("AccessKey", config.BUNNYNET_WRITE_API_KEY)
	req.Header.Add("content-type", "application/octet-stream")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

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
