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
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		us.log.Errorf("controller.user.CreateUser.Decode couldn't decode user, %s", err)
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
		return
	}

	err = us.userService.CreateUser(user)
	if err != nil {
		us.log.Errorf("controller.user.CreateUser couldn't create user, %s", err)
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write([]byte(`{"message": "The user has been created"}`))
	if err != nil {
		us.log.Errorf("controller.user.CreateUser, %s", err)
		return
	}
}

func (us UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uuid, found := mux.Vars(r)["id"]
	if !found {
		us.log.Info(": [INFO] Id not found ")
		w.WriteHeader((http.StatusInternalServerError))
		return
	}
	user, err := us.userService.GetUser(uuid)
	if err != nil {
		us.log.Errorf("controller.user.GetUser couldn't get user, %s", err)
		w.WriteHeader((http.StatusInternalServerError))

		return
	}

	byteUser, err := json.Marshal(user)
	if err != nil {
		us.log.Errorf("controller.user.GetUser couldn't parse User")
		w.WriteHeader((http.StatusInternalServerError))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write(byteUser)
	if err != nil {
		us.log.Errorf("controller.user.GetUser, %s", err)
		return
	}
}

func (us UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uuid, found := mux.Vars(r)["id"]
	if !found {
		us.log.Info(": [INFO] Id not found ")
		w.WriteHeader((http.StatusInternalServerError))
		return
	}

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		us.log.Errorf("controller.user.UpdateUser.Decode couldn't decode user, %s", err)
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
		return
	}

	err = us.userService.UpdateUser(uuid, user)
	if err != nil {
		us.log.Errorf("controller.user.UpdateUser couldn't update user, %s", err)
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write([]byte(`{"message": "The user has been changed"}`))
	if err != nil {
		us.log.Errorf("controller.user.UpdateUser, %s", err)
		return
	}
}

func (us UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uuid, found := mux.Vars(r)["id"]
	if !found {
		us.log.Info(": [INFO] Id not found ")
		w.WriteHeader((http.StatusInternalServerError))
		return
	}

	err := us.userService.DeleteUser(uuid)
	if err != nil {
		us.log.Errorf("controller.user.DeleteUser couldn't delete user, %s", err)
		w.WriteHeader((http.StatusInternalServerError))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write([]byte(`{"message": "The user has been deleted"}`))
	if err != nil {
		us.log.Errorf("controller.user.DeleteUser, %s", err)
		return
	}
}

func (us UserController) GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := us.userService.GetUserList()
	if err != nil {
		us.log.Errorf("controller.user.GetUserList couldn't get user list, %s", err)
		w.WriteHeader((http.StatusInternalServerError))
		return
	}

	byteUsers, err := json.Marshal(users)
	if err != nil {
		us.log.Errorf("controller.user.GetUser couldn't parse User")
		w.WriteHeader((http.StatusInternalServerError))
		return
	}
	w.WriteHeader((http.StatusOK))
	_, err = w.Write(byteUsers)
	if err != nil {
		us.log.Errorf("controller.user.GetUserList, %s", err)
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
