package main

import (
	"log"

	octobus "github.com/ghatdev/octobus-go"
)

func main() {
	logger := octobus.New(octobus.DefaultConfig("", "Default"))
	err := logger.Dial()
	if err != nil {
		log.Println(err)
	}

	logger.Log("a", "DEBUG", "E", "p")
}
