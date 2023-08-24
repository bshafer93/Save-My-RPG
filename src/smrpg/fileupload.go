package smrpg

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"fmt"
)

type SaveFileChunk struct {
	TotalChunks uint32 //4 bytes
	ChunkNumber uint32 //4 bytes
	ChunkSize   uint32 //4 bytes
	Data        []byte
} //512012 bytes

type FullSaveFileJson struct {
	FolderName string `json:"FolderName"`
	FolderSize int64  `json:"FolderSize"`
	Data       string `json:"ZipDataB64"`
}

func SerializeChunk(c *SaveFileChunk) []byte {

	buf := make([]byte, 512012)
	binary.LittleEndian.PutUint32(buf[0:], c.TotalChunks)
	binary.LittleEndian.PutUint32(buf[4:], c.ChunkNumber)
	binary.LittleEndian.PutUint32(buf[8:], c.ChunkSize)
	buf = append(buf[:12], c.Data...)
	return buf
}

func DeserializeChunk(chunk []byte) *SaveFileChunk {
	sfc := SaveFileChunk{}
	sfc.TotalChunks = binary.LittleEndian.Uint32(chunk[0:])
	sfc.ChunkNumber = binary.LittleEndian.Uint32(chunk[4:])
	sfc.ChunkSize = binary.LittleEndian.Uint32(chunk[8:])
	sfc.Data = bytes.Clone(chunk[12:])
	return &sfc
}

func EstablishFileLink() {
	_, err := LoadConfiguration("/go/src/savemyrpgserver/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	cer, err := tls.LoadX509KeyPair(config.SERVER_CERT, config.SERVER_KEY)
	if err != nil {
		fmt.Println(err)
		return
	}
	cert_config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":"+config.SERVER_PORT, cert_config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	/*
		for {
			conn, err := ln.Accept()
			fmt.Println("Connected Received!")
			if err != nil {
				log.Println(err)
				continue
			}
			//go HandleSendFile(conn)
		}
	*/
}
