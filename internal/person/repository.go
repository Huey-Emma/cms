package person

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/huey-emma/cms/internal/utils/lib"
)

var ErrResourceNotFound = errors.New("resource not found")

type DBTX interface {
	QueryRowContext(context.Context, string, ...any) *sql.Row
	ExecContext(context.Context, string, ...any) (sql.Result, error)
}

type Repository interface {
	InsertPerson(context.Context, []byte) (lib.Map[any], error)
	FindPerson(context.Context, int) (lib.Map[any], error)
	UpdatePerson(context.Context, lib.Map[any]) error
	DeletePerson(context.Context, int) error
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db}
}

func (r *repository) InsertPerson(ctx context.Context, payload []byte) (lib.Map[any], error) {
	q := `INSERT INTO persons (info) VALUES ($1) RETURNING id;`
	row := r.db.QueryRowContext(ctx, q, payload)

	var id int

	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	out := make(lib.Map[any])

	if err := json.Unmarshal(payload, &out); err != nil {
		return nil, err
	}

	out["id"] = id

	return out, nil
}

func (r *repository) FindPerson(ctx context.Context, id int) (lib.Map[any], error) {
	q := `SELECT * FROM persons WHERE id = $1;`
	row := r.db.QueryRowContext(ctx, q, id)

	result := new(struct {
		ID   int
		Info []byte
	})

	err := row.Scan(&result.ID, &result.Info)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrResourceNotFound
	}

	if err != nil {
		return nil, err
	}

	out := make(lib.Map[any])

	out["id"] = result.ID

	info := make(lib.Map[any])

	err = json.Unmarshal(result.Info, &info)
	if err != nil {
		return nil, err
	}

	for k, v := range info {
		out[k] = v
	}

	return out, nil
}

func (r *repository) UpdatePerson(ctx context.Context, person lib.Map[any]) error {
	q := `UPDATE persons SET info = $1 WHERE id = $2;`

	id := person["id"]
	delete(person, "id")

	info, err := json.Marshal(person)
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, q, info, id)

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrResourceNotFound
	}

	return nil
}

func (r *repository) DeletePerson(ctx context.Context, id int) error {
	q := `DELETE FROM persons WHERE id = $1;`

	result, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrResourceNotFound
	}

	return nil
}
