package main

import (
	"github.com/joho/godotenv"
	"onetomany/databaseconnection"
	"onetomany/handler"
)

func main() {
	godotenv.Load("config.env")
	databaseconnection.Connection()
	handler.Handler()

}
