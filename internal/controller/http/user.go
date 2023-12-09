package http

import (
	"github.com/gorilla/mux"
	"net/http"
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

}

func (u *userController) removeUser(w http.ResponseWriter, r *http.Request) {

}
