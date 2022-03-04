package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

var Blockchain []Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		dummyBlock := Block{0, t.String(), "", "", ""}
		spew.Dump(dummyBlock)

		Blockchain = append(Blockchain, dummyBlock)
	}()

	log.Fatal(run())
}
