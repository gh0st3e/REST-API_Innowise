package main

import (
	"InnowisePreTraineeTask/internal/controller"
	"InnowisePreTraineeTask/internal/database"
	"InnowisePreTraineeTask/internal/repository"
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
	userServer := controller.NewUserServer(log, userService)

	err := userRepository.CheckTable()
	if err != nil {
		log.Errorf("Couldn't create table")
	}

	r := mux.NewRouter()

	userServer.Mount(r)

	log.Info("server is listening on 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
