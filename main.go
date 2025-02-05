package main

import (
	"odkt/server"
	"odkt/server/db"
	"odkt/server/helper"
	"os"
)

func main() {
	os.MkdirAll("./data", 0750)
	db.InitDatabase()
	helper.ImportCardsToDB()
	server.Start()
}
