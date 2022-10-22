package server

import (
	"InnowisePreTraineeTask/internal/entity"
	"InnowisePreTraineeTask/internal/util"
	"encoding/json"
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
		log.Println("server.user.CreateUser couldn't create user, %s", err)
		return
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
	user, err := us.userService.GetUser(uuid)
	if err != nil {
		log.Println("server.user.GetUser couldn't get user, %s", err)
		return
	}

	byteUser, err := json.Marshal(user)
	if err != nil {
		log.Println("server.user.GetUser couldn't parse User")
		return
	}

	w.Write(byteUser)
}

func (us UserServer) UpdateUser(w http.ResponseWriter, r *http.Request) {
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
		log.Println("server.user.UpdateUser couldn't update user, %s", err)
		return
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
		log.Println("server.user.DeleteUser couldn't delete user, %s", err)
		return
	}

	w.Write([]byte(`{"message": "The user has been deleted"}`))
}

func (us UserServer) GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	users, err := us.userService.GetUserList()
	if err != nil {
		log.Println("server.user.GetUserList couldn't get user list, %s", err)
		return
	}

	byteUsers, err := json.Marshal(users)
	if err != nil {
		log.Println("server.user.GetUser couldn't parse User")
		return
	}

	w.Write(byteUsers)
}

func (server *UserServer) Mount(r *mux.Router) {
	r.HandleFunc(util.PostUser, server.CreateUser).Methods("POST")
	r.HandleFunc(util.GetUser, server.GetUser).Methods("GET")
	r.HandleFunc(util.PutUser, server.UpdateUser).Methods("PUT")
	r.HandleFunc(util.DelUser, server.DeleteUser).Methods("DELETE")
	r.HandleFunc(util.GetUserList, server.GetUserList).Methods("GET")
}
