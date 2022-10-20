package main

import (
	"InnowisePreTraineeTask/internal/database"
	"InnowisePreTraineeTask/internal/server"
)

func main() {

	db := database.Connect()
	server.InitServer(db)
}
