package main

import (
	"bytes"
	"encoding/binary"
	"savemyrpg/smrpg"
	"testing"
)

func TestDeserializeChunk(t *testing.T) {
	c := smrpg.SaveFileChunk{}
	c.TotalChunks = 789
	c.ChunkNumber = 89
	c.ChunkSize = 512000
	c.Data = bytes.Repeat([]byte("A"), 512000)
	buf := smrpg.SerializeChunk(&c)

	c_deserialized := smrpg.DeserializeChunk(buf)

	t.Logf("Total Chunks: %d", c_deserialized.TotalChunks)
	t.Logf("Chunk Number: %d", c_deserialized.ChunkNumber)
	t.Logf("Chunk Size %d", c_deserialized.ChunkSize)

	if c_deserialized.TotalChunks != c.TotalChunks {
		t.Errorf("Total Chunks Expected: %v != %v", c.TotalChunks, c_deserialized.TotalChunks)
	}

	if c_deserialized.ChunkNumber != c.ChunkNumber {
		t.Errorf("Chunk Number Expected: %v != %v", c.ChunkNumber, c_deserialized.ChunkNumber)
	}

	if c_deserialized.ChunkSize != c.ChunkSize {
		t.Errorf("Chunk Size Expected: %v != %v", c.ChunkSize, c_deserialized.ChunkSize)
	}

	if !bytes.Equal(c_deserialized.Data, c.Data) {
		t.Error("Chunk Data not equal!")
	}
}

func TestSerializeChunk(t *testing.T) {
	c := smrpg.SaveFileChunk{}
	c.TotalChunks = 10
	c.ChunkNumber = 7
	c.ChunkSize = 512000
	c.Data = bytes.Repeat([]byte("A"), 512000)
	buf := smrpg.SerializeChunk(&c)
	var tn uint32 = binary.LittleEndian.Uint32(buf[0:])
	var cn uint32 = binary.LittleEndian.Uint32(buf[4:])
	var cs uint32 = binary.LittleEndian.Uint32(buf[8:])

	t.Logf("Total Chunks: %d", tn)
	t.Logf("Chunk Number: %d", cn)
	t.Logf("Chunk Size %d", cs)

	if tn != c.TotalChunks {
		t.Errorf("Total Chunks Expected: %v != %v", c.TotalChunks, tn)
	}

	if cn != c.ChunkNumber {
		t.Errorf("Chunk Number Expected: %v != %v", c.ChunkNumber, cn)
	}

	if cs != c.ChunkSize {
		t.Errorf("Chunk Size Expected: %v != %v", c.ChunkSize, cs)
	}

	if !bytes.Equal(buf[12:], c.Data) {
		t.Error("Chunk Data not equal!")
	}
}

func TestInitBunny(t *testing.T) {
	_, err := smrpg.LoadConfiguration("config.json")
	if err != nil {
		t.Error(err)
	}
	smrpg.InitBunny()
	smrpg.BunnyListAllFIles()

}

func TestBunnyDownloadFile(t *testing.T) {
	_, err := smrpg.LoadConfiguration("config.json")
	if err != nil {
		t.Error(err)
	}
	smrpg.InitBunny()
	smrpg.BunnyDownloadFile("Norbertle-3812316722__TheHollow.zip")

}

func TestBunnyUploadFile(t *testing.T) {
	_, err := smrpg.LoadConfiguration("config.json")
	if err != nil {
		t.Error(err)
	}
	smrpg.InitBunny()
	smrpg.BunnyUploadFile("TempFolder", "AnotherZipBytesTheDust.zip", `Norbertle-3812316722__Poop.zip`, `C:\Users\brent\Documents\Programming\SaveMyRPGServer\src\Norbertle-3812316722__Poop.zip`)

}
