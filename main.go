package main

import (
	"world-api/server"
)

func main() {
	server := server.NewServer()

	server.Run()
}
