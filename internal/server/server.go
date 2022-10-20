package server

import (
	"InnowisePreTraineeTask/internal/repository"
	"InnowisePreTraineeTask/internal/service"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write([]byte(`{"message": "Main page"}`))
}

type IServer interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	sv service.IService
}

func New(repo service.IService) IServer {
	return &Server{repo}
}

// Создать юзера
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	//TODO так вообще можно делать, типо вот так передвать аргумент?)
	err := service.CreateUser(r.Body)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(`{"message": "The user has been created"}`))
}

// Получить юзера по айди
func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	uuid, found := mux.Vars(r)["id"]
	if !found {
		log.Println(": [INFO] Id not found ")
		return
	}

	user, _ := s.sv.GetUser(uuid)

	w.Write(user)

}

// Апдейтнуть юзера
func EditUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	uuid, found := mux.Vars(r)["id"]
	if !found {
		log.Println(": [INFO] Id not found ")
		return
	}
	err := service.UpdateUser(uuid, r.Body)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(`{"message": "The user has been changed"}`))
}

// Удалить юзера
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	uuid, found := mux.Vars(r)["id"]
	if !found {
		log.Println(": [INFO] Id not found ")
		return
	}
	err := service.DeleteUser(uuid)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(`{"message": "The user has been deleted"}`))
}

// Получить лист юзеров
func GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))

	users, err := service.GetUserList()
	if err != nil {
		panic(err)
	}

	w.Write(users)
}

func InitServer(db *sql.DB) {
	r := mux.NewRouter()

	rep := repository.NewRepo(db)
	store := service.New(rep)
	serv := New(store)

	r.HandleFunc("/", ServeHTTP)
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", serv.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", EditUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", GetUserList).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
