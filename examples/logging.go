package main

import (
	"log"

	octobus "github.com/ghatdev/octobus-go"
)

func main() {
	logger := octobus.New(octobus.DefaultConfig("", "test1"))
	err := logger.Dial()
	if err != nil {
		log.Println(err)
	}

	logger.Log("a", "b", "E", "p")
}
