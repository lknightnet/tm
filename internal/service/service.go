package service

import (
	"context"
	"task-manager/internal/model"
	"task-manager/internal/repository"
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

type Services struct {
	MP MethodProject
	MU MethodUser
	MN MethodNote
}

type Dependencies struct {
	Repos *repository.Repositories
}

func NewServices(deps *Dependencies) *Services {
	return &Services{
		MN: nil,
		MP: newProjectService(deps.Repos.MP),
		MU: nil,
	}
}
