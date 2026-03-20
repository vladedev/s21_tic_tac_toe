package register

import (
	"Project03-Go_Bootcamp/internal/domain"
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Get(ctx context.Context, id string) (domain.Game, error)
	Save(ctx context.Context, g domain.Game, id string) error
}

type ServiceRegistr struct {
	repo Repository
}

func New(repo Repository) *ServiceRegistr {
	return &ServiceRegistr{repo: repo}
}

func (s *ServiceRegistr) RegisterGame(ctx context.Context) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	idStr := id.String()

	if err = s.repo.Save(ctx, domain.NewGame(idStr), idStr); err != nil {
		return "", err
	}
	return idStr, nil
}
