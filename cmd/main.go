package main

import (
	"log"
	"main/internal/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Println(err)
		return
	}
}
