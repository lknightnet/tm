package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"task-manager/internal/model"
	"task-manager/internal/service"
)

type userController struct {
	MU service.MethodUser
}

func newUserController(mu service.MethodUser) *userController {
	return &userController{MU: mu}
}

func (u *userController) newUserRoutes(r *mux.Router) {
	r.HandleFunc("/user/add", u.addUser).Methods(http.MethodPost)
	r.HandleFunc("/user/remove", u.removeUser).Methods(http.MethodPost)
}

func (u *userController) addUser(w http.ResponseWriter, r *http.Request) {
	var body model.ProjectUser

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		return
	}

	err = u.MU.AddUser(r.Context(), &body)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u *userController) removeUser(w http.ResponseWriter, r *http.Request) {
	var body model.ProjectUser

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		return
	}

	err = u.MU.RemoveUser(r.Context(), &body)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
