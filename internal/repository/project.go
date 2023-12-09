package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"log"
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

	log.Println(sql)
	log.Println(args)

	_, err = p.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "CreateProject: fail to exec")
	}
	return nil
}

func (p *projectRepository) RemoveProject(ctx context.Context, proj *model.Project) error {
	sql, args, err := p.pg.Builder.Delete("projects").Where(squirrel.Eq{"id": proj.ID}).ToSql()
	if err != nil {
		return errors.Wrap(err, "RemoveProject: sql generator is failed")
	}

	log.Println(sql)
	log.Println(args)

	_, err = p.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "RemoveProject: fail to exec")
	}
	return nil
}

func newProjectRepository(pg *database.Postgres) *projectRepository {
	return &projectRepository{pg: pg}
}
