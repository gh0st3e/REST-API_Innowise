package main

import (
	"InnowisePreTraineeTask/internal/database"
	"InnowisePreTraineeTask/internal/repository"
	"InnowisePreTraineeTask/internal/server"
	"InnowisePreTraineeTask/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log := logrus.New()
	db := database.Connect()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(log, userRepository)
	userServer := server.NewUserServer(userService)

	r := mux.NewRouter()

	userServer.Mount(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
