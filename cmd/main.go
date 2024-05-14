package main

import (
	app "bashscripts/internal"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

//можно запустить всё командой make в консоли
func main() {
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
