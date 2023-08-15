package main

// import the package we need to use
import (
	"log"
	"savemyrpg/smrpg"

	_ "github.com/lib/pq"
)

func main() {
	smrpg.Init()

	log.Fatal(smrpg.Start())

}
