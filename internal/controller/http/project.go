package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"task-manager/internal/model"
	"task-manager/internal/service"
)

type projectController struct {
	MP service.MethodProject
}

func newProjectController(mp service.MethodProject) *projectController {
	return &projectController{MP: mp}
}

func (p *projectController) newProjectRoutes(r *mux.Router) {
	r.HandleFunc("/project/create", p.createProject).Methods(http.MethodPost)
	r.HandleFunc("/project/remove", p.removeProject).Methods(http.MethodPost)
}

func (p *projectController) createProject(w http.ResponseWriter, r *http.Request) {
	var body model.Project

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		return
	}

	err = p.MP.CreateProject(r.Context(), &body)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *projectController) removeProject(w http.ResponseWriter, r *http.Request) {
	var body model.Project
	log.Println(body)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		return
	}

	err = p.MP.RemoveProject(r.Context(), &body)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
