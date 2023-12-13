package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"task-manager/internal/model"
	"task-manager/pkg/database"
)

type userRepository struct {
	pg *database.Postgres
}

func (u *userRepository) AddUser(ctx context.Context, user *model.ProjectUser) error {
	sql, args, err := u.pg.Builder.Insert("projectusers").
		Columns("iduser", "idproject").
		Values(user.IDUser, user.ID).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "AddUser: sql generator is failed")
	}

	_, err = u.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "AddUser: fail to exec")
	}
	return nil
}

func (u *userRepository) RemoveUser(ctx context.Context, user *model.ProjectUser) error {
	sql, args, err := u.pg.Builder.Delete("projectusers").Where(squirrel.Eq{"idproject": user.ID, "iduser": user.IDUser}).ToSql()
	if err != nil {
		return errors.Wrap(err, "RemoveUser: sql generator is failed")
	}

	_, err = u.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "RemoveUser: fail to exec")
	}
	return nil
}

func newUserRepository(pg *database.Postgres) *userRepository {
	return &userRepository{pg: pg}
}

var _ MethodUser = (*userRepository)(nil)
