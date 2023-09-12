package person

import (
	"context"
	"encoding/json"
	"time"

	"github.com/huey-emma/cms/internal/utils/lib"
)

type Service interface {
	InsertPerson(context.Context, map[string]string) (lib.Map[any], error)
	FindPerson(context.Context, int) (lib.Map[any], error)
	UpdatePerson(context.Context, lib.Map[any]) (lib.Map[any], error)
	DeletePerson(context.Context, int) error
}

type service struct {
	repository Repository
	timeout    time.Duration
}

func NewService(repository Repository) Service {
	timeout := 3 * time.Second
	return &service{repository, timeout}
}

func (s *service) InsertPerson(
	ctx context.Context,
	payload map[string]string,
) (lib.Map[any], error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return s.repository.InsertPerson(ctx, b)
}

func (s *service) FindPerson(ctx context.Context, id int) (lib.Map[any], error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	return s.repository.FindPerson(ctx, id)
}

func (s *service) UpdatePerson(ctx context.Context, person lib.Map[any]) (lib.Map[any], error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.repository.UpdatePerson(ctx, person); err != nil {
		return person, err
	}

	return person, nil
}

func (s *service) DeletePerson(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	return s.repository.DeletePerson(ctx, id)
}
