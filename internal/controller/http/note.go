package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"task-manager/internal/model"
	"task-manager/internal/service"
)

type noteController struct {
	MN service.MethodNote
}

func newNoteController(mn service.MethodNote) *noteController {
	return &noteController{MN: mn}
}

func (n *noteController) newNoteRoutes(r *mux.Router) {
	r.HandleFunc("/note/create", n.createNote).Methods(http.MethodPost)
	r.HandleFunc("/note/remove", n.removeNote).Methods(http.MethodPost)
	r.HandleFunc("/note/update", n.updateNote).Methods(http.MethodPost)
}

func (n *noteController) createNote(w http.ResponseWriter, r *http.Request) {
	var body model.Note

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		return
	}

	err = n.MN.CreateNote(r.Context(), &body)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (n *noteController) removeNote(w http.ResponseWriter, r *http.Request) {
	var body model.Note

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		return
	}

	err = n.MN.RemoveNote(r.Context(), &body)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (n *noteController) updateNote(w http.ResponseWriter, r *http.Request) {
	var body model.Note

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		return
	}

	err = n.MN.UpdateCompletenessNote(r.Context(), &body)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
