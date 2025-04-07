package main

import (
	"IPfinder/app"
	"log"
	"os"
)

func main() {
	aplication := app.Init()

	failure := aplication.Run(os.Args)
	if failure != nil {
		log.Fatal(failure)
	}
}
