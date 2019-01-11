package main

import (
	"log"

	octobus "github.com/ghatdev/octobus-go"
)

func main() {
	logger := octobus.New(&octobus.BasicLogger{Config: octobus.DefaultConfig("localhost:17000", "default")})
	err := logger.Dial()
	if err != nil {
		log.Fatalln(err)
	}

	logger.Log("b", "E", "p")
	logger.Debug("d", "blabla")
}
