package main

import (
	"log"

	"github.com/ahmdrz/divan-e-shams/application/server"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
