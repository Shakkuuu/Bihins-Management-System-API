package main

import (
	db "Goods-Management-System-API/dbgo"
	"Goods-Management-System-API/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
