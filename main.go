package main

import (
	db "Bihins-Management-System-API/dbgo"
	"Bihins-Management-System-API/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
