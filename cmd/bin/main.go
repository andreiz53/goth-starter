package main

import (
	"goth/server"
)

func main() {
	server := server.NewServer(":8080")
	server.Init()
}
