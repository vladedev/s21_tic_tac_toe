package datasource

import (
	"Project03-Go_Bootcamp/internal/domain"
	"context"
	"errors"
	"sync"
)

type MainRepo struct {
	data sync.Map
}

func New() *MainRepo {
	return &MainRepo{}
}

func (m *MainRepo) Save(ctx context.Context, g domain.Game, id string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		m.data.Store(id, ToDS(g))
		return nil
	}
}

func (m *MainRepo) Get(ctx context.Context, id string) (domain.Game, error) {
	select {
	case <-ctx.Done():
		return domain.Game{}, ctx.Err()
	default:
		value, ok := m.data.Load(id)
		if !ok {
			return domain.Game{}, errors.New("game not found: " + id)
		}
		return ToDomain(value.(GameDS)), nil
	}
}
