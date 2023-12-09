package service

import (
	"context"
	"github.com/google/uuid"
	"task-manager/internal/model"
	"task-manager/internal/repository"
)

type projectService struct {
	mp repository.MethodProject
}

func (p *projectService) CreateProject(ctx context.Context, proj *model.Project) error {
	proj.ID = uuid.NewString()
	return p.mp.CreateProject(ctx, proj)
}

func (p *projectService) RemoveProject(ctx context.Context, proj *model.Project) error {
	return p.mp.RemoveProject(ctx, proj)
}

func newProjectService(mp repository.MethodProject) *projectService {
	return &projectService{mp: mp}
}

var _ MethodProject = (*projectService)(nil)
