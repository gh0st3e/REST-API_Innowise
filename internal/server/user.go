package server

import (
	"InnowisePreTraineeTask/internal/entity"
	"InnowisePreTraineeTask/internal/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (us UserServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	var user entity.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := us.userService.CreateUser(user)
	if err != nil {
		panic(err)
	}
	
	w.Write([]byte(`{"message": "The user has been created"}`))
}

func (us UserServer) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	uuid, found := mux.Vars(r)["id"]
	if !found {
		log.Println(": [INFO] Id not found ")
		return
	}
	user, _ := us.userService.GetUser(uuid)

	byteUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(byteUser)
}

func (us UserServer) EditUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	uuid, found := mux.Vars(r)["id"]
	if !found {
		log.Println(": [INFO] Id not found ")
		return
	}

	var user entity.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := us.userService.UpdateUser(uuid, user)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(`{"message": "The user has been changed"}`))
}

func (us UserServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	uuid, found := mux.Vars(r)["id"]
	if !found {
		log.Println(": [INFO] Id not found ")
		return
	}

	err := us.userService.DeleteUser(uuid)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(`{"message": "The user has been deleted"}`))
}

func (us UserServer) GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	users, err := us.userService.GetUserList()
	if err != nil {
		panic(err)
	}

	byteUsers, err := json.Marshal(users)
	if err != nil {

	}

	w.Write(byteUsers)
}

func (server *UserServer) Mount(r *mux.Router) {
	r.HandleFunc(util.PostUser, server.CreateUser).Methods("POST")
	r.HandleFunc(util.GetUser, server.GetUser).Methods("GET")
	r.HandleFunc(util.PutUser, server.EditUser).Methods("PUT")
	r.HandleFunc(util.DelUser, server.DeleteUser).Methods("DELETE")
	r.HandleFunc(util.GetUserList, server.GetUserList).Methods("GET")
}
