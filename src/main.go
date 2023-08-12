package main

// import the package we need to use
import (
	"log"
	"savemyrpg/smrpg"

	_ "github.com/lib/pq"
)

func main() {
	smrpg.Init()
	smrpg.ZipSaveFile("Norbertle-31112316728_smrpg")

	log.Fatal(smrpg.Start())

}
