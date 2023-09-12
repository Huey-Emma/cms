package person

import (
	"context"
	"database/sql"

	"github.com/huey-emma/cms/internal/utils/lib"
)

type DBTX interface {
	QueryRowContext(context.Context, string, ...any) *sql.Row
	ExecContext(context.Context, string, ...any) (sql.Result, error)
}

type Repository interface {
	InsertPerson(context.Context, []byte) (lib.Map[any], error)
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db}
}

func (r *repository) InsertPerson(ctx context.Context, payload []byte) (lib.Map[any], error) {
	return nil, nil
}
