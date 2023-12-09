package repository

import (
	"context"
	"task-manager/internal/model"
	"task-manager/pkg/database"
)

type MethodProject interface {
	CreateProject(ctx context.Context, proj *model.Project) error
	RemoveProject(ctx context.Context, proj *model.Project) error
}

type MethodUser interface {
	AddUser(ctx context.Context, user *model.ProjectUser) error
	RemoveUser(ctx context.Context, user *model.ProjectUser) error
}

type MethodNote interface {
	CreateNote(ctx context.Context, note *model.Note) error
	UpdateCompletenessNote(ctx context.Context, note *model.Note) error
	RemoveNote(ctx context.Context, note *model.Note) error
}

type Repositories struct {
	MP MethodProject
	MU MethodUser
	MN MethodNote
}

func NewRepositories(pg *database.Postgres) *Repositories {
	return &Repositories{
		MP: newProjectRepository(pg),
		MU: nil,
		MN: nil,
	}
}
