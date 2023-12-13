package service

import (
	"context"
	"task-manager/internal/model"
	"task-manager/internal/repository"
)

type userService struct {
	mu repository.MethodUser
}

func (u *userService) AddUser(ctx context.Context, user *model.ProjectUser) error {
	return u.mu.AddUser(ctx, user)
}

func (u *userService) RemoveUser(ctx context.Context, user *model.ProjectUser) error {
	return u.mu.RemoveUser(ctx, user)
}

func newUserService(mu repository.MethodUser) *userService {
	return &userService{mu: mu}
}

var _ MethodUser = (*userService)(nil)
