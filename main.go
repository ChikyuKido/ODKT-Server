package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"odkt/server"
	"odkt/server/db"
	"odkt/server/helper"
	"odkt/server/store"
	"os"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			logrus.Errorf("Error loading .env file")
		}
	}
	os.MkdirAll("./data", 0750)
	store.InitStores()
	db.InitDatabase()
	helper.ImportCardsToDB()
	server.Start()
}
