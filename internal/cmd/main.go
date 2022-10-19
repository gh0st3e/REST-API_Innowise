package main

import (
	"InnowisePreTraineeTask/internal/database"
	"InnowisePreTraineeTask/internal/server"
)

func main() {
	server.InitServer()
	database.Connect()
}
