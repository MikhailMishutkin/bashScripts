package main

import (
	app "bashscripts/internal"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
