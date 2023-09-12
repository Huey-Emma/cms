package person

import (
	"context"

	"github.com/huey-emma/cms/internal/utils/lib"
)

type Service interface {
	InsertPerson(context.Context, map[string]string) (lib.Map[any], error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) InsertPerson(
	ctx context.Context,
	payload map[string]string,
) (lib.Map[any], error) {
	return nil, nil
}
