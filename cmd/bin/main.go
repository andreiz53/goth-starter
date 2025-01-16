package main

import (
	"goth/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	server := server.NewServer(os.Getenv("PORT"))
	server.Init()
}
