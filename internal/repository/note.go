package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"task-manager/internal/model"
	"task-manager/pkg/database"
)

type noteRepository struct {
	pg *database.Postgres
}

func (n *noteRepository) CreateNote(ctx context.Context, note *model.Note) error {
	sql, args, err := n.pg.Builder.Insert("notes").
		Columns("id", "idproject", "description", "completeness").
		Values(note.ID, note.IDProject, note.DescriptionOfNote, note.Completeness).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "CreateNote: sql generator is failed")
	}

	_, err = n.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "CreateNote: fail to exec")
	}
	return nil
}

func (n *noteRepository) UpdateCompletenessNote(ctx context.Context, note *model.Note) error {
	sql, args, err := n.pg.Builder.Update("notes").Set("completeness", note.Completeness).
		Where(squirrel.Eq{"id": note.ID}).ToSql()
	if err != nil {
		return errors.Wrap(err, "UpdateCompletenessNote: sql generator is failed")
	}

	_, err = n.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "UpdateCompletenessNote: fail to exec")
	}
	return nil
}

func (n *noteRepository) RemoveNote(ctx context.Context, note *model.Note) error {
	sql, args, err := n.pg.Builder.Delete("notes").Where(squirrel.Eq{"id": note.ID}).ToSql()
	if err != nil {
		return errors.Wrap(err, "RemoveNote: sql generator is failed")
	}

	_, err = n.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "RemoveNote: fail to exec")
	}
	return nil
}

func newNoteRepository(pg *database.Postgres) *noteRepository {
	return &noteRepository{pg: pg}
}
