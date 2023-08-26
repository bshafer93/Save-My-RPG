package smrpg

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

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
