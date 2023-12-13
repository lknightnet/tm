package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"task-manager/internal/model"
	"task-manager/pkg/database"
)

type projectRepository struct {
	pg *database.Postgres
}

func (p *projectRepository) CreateProject(ctx context.Context, proj *model.Project) error {
	sql, args, err := p.pg.Builder.Insert("projects").
		Columns("iduser", "id", "name", "description").
		Values(proj.IDUser, proj.ID, proj.Name, proj.DescriptionOfProject).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "CreateProject: sql generator is failed")
	}

	_, err = p.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "CreateProject: fail to exec")
	}
	return nil
}

func (p *projectRepository) RemoveProject(ctx context.Context, proj *model.Project) error {
	sql, args, err := p.pg.Builder.Delete("notes").Where(squirrel.Eq{"idproject": proj.ID}).ToSql()

	if err != nil {
		return errors.Wrap(err, "RemoveProject: sql delete notes generator is failed")
	}

	_, err = p.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "RemoveProject: fail to exec delete notes")
	}

	sql, args, err = p.pg.Builder.Delete("projectusers").Where(squirrel.Eq{"idproject": proj.ID}).ToSql()

	if err != nil {
		return errors.Wrap(err, "RemoveProject: sql delete users generator is failed")
	}

	_, err = p.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "RemoveProject: fail to exec delete users")
	}

	sql, args, err = p.pg.Builder.Delete("projects").Where(squirrel.Eq{"id": proj.ID}).ToSql()
	if err != nil {
		return errors.Wrap(err, "RemoveProject: sql delete projects generator is failed")
	}

	_, err = p.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "RemoveProject: fail to exec delete projects")
	}

	return nil
}

func newProjectRepository(pg *database.Postgres) *projectRepository {
	return &projectRepository{pg: pg}
}
