package service

import (
	"context"
	"github.com/google/uuid"
	"task-manager/internal/model"
	"task-manager/internal/repository"
)

type noteService struct {
	mn repository.MethodNote
}

func (n *noteService) CreateNote(ctx context.Context, note *model.Note) error {
	note.ID = uuid.NewString()
	return n.mn.CreateNote(ctx, note)
}

func (n *noteService) UpdateCompletenessNote(ctx context.Context, note *model.Note) error {
	return n.mn.UpdateCompletenessNote(ctx, note)
}

func (n *noteService) RemoveNote(ctx context.Context, note *model.Note) error {
	return n.mn.RemoveNote(ctx, note)
}

func newNoteService(mn repository.MethodNote) *noteService {
	return &noteService{mn: mn}
}

var _ MethodNote = (*noteService)(nil)
