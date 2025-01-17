package main

import (
	"goth/cookies"
	"goth/server"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cookieStore := cookies.NewCookieStore()
	server := server.NewServer(os.Getenv("PORT"), cookieStore)
	server.Init()
}
