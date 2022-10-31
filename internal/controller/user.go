package controller

import (
	"InnowisePreTraineeTask/internal/entity"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (us UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user entity.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := us.userService.CreateUser(user)
	if err != nil {
		us.log.Printf("controller.user.CreateUser couldn't create user, %s", err)
		w.WriteHeader((http.StatusNotFound))
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err)))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write([]byte(`{"message": "The user has been created"}`))
	if err != nil {
		us.log.Printf("controller.user.CreateUser, %s", err)
		return
	}
}

func (us UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uuid, found := mux.Vars(r)["id"]
	if !found {
		us.log.Println(": [INFO] Id not found ")
		w.WriteHeader((http.StatusNotFound))
		return
	}
	user, err := us.userService.GetUser(uuid)
	if err != nil {
		us.log.Printf("controller.user.GetUser couldn't get user, %s", err)
		w.WriteHeader((http.StatusNotFound))

		return
	}

	byteUser, err := json.Marshal(user)
	if err != nil {
		us.log.Printf("controller.user.GetUser couldn't parse User")
		w.WriteHeader((http.StatusNotFound))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write(byteUser)
	if err != nil {
		us.log.Printf("controller.user.GetUser, %s", err)
		return
	}
}

func (us UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uuid, found := mux.Vars(r)["id"]
	if !found {
		us.log.Println(": [INFO] Id not found ")
		w.WriteHeader((http.StatusNotFound))
		return
	}

	var user entity.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := us.userService.UpdateUser(uuid, user)
	if err != nil {
		us.log.Printf("controller.user.UpdateUser couldn't update user, %s", err)
		w.WriteHeader((http.StatusNotFound))
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err)))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write([]byte(`{"message": "The user has been changed"}`))
	if err != nil {
		us.log.Printf("controller.user.UpdateUser, %s", err)
		return
	}
}

func (us UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uuid, found := mux.Vars(r)["id"]
	if !found {
		us.log.Println(": [INFO] Id not found ")
		w.WriteHeader((http.StatusNotFound))
		return
	}

	err := us.userService.DeleteUser(uuid)
	if err != nil {
		us.log.Printf("controller.user.DeleteUser couldn't delete user, %s", err)
		w.WriteHeader((http.StatusNotFound))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write([]byte(`{"message": "The user has been deleted"}`))
	if err != nil {
		us.log.Printf("controller.user.DeleteUser, %s", err)
		return
	}
}

func (us UserController) GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := us.userService.GetUserList()
	if err != nil {
		us.log.Printf("controller.user.GetUserList couldn't get user list, %s", err)
		w.WriteHeader((http.StatusNotFound))
		return
	}

	byteUsers, err := json.Marshal(users)
	if err != nil {
		us.log.Println("controller.user.GetUser couldn't parse User")
		w.WriteHeader((http.StatusNotFound))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write(byteUsers)
	if err != nil {
		us.log.Printf("controller.user.GetUserList, %s", err)
		return
	}
}

func (c *UserController) Mount(r *mux.Router) {
	r.HandleFunc("/users", c.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", c.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", c.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", c.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", c.GetUserList).Methods("GET")
}
