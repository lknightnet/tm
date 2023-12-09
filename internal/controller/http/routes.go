package http

import (
	"github.com/gorilla/mux"
	"task-manager/internal/service"
)

func NewRoutes(srv *service.Services, r *mux.Router) {
	ctrlProject := newProjectController(srv.MP)
	ctrlProject.newProjectRoutes(r)

	ctrlNote := newNoteController(srv.MN)
	ctrlNote.newNoteRoutes(r)

	//ctrlUser := newUserController(srv.MU)
	//ctrlUser.newUserRoutes(r)
}
