package repository

import (
	"caffinity/internal/db"
	"context"
)

type CafeRepository struct {
	Queries *db.Queries
}

func NewCafeRepository(q *db.Queries) *CafeRepository {
	return &CafeRepository{Queries: q}
}

func (r *CafeRepository) GetAll(ctx context.Context) ([]db.Cafe, error) {
	return r.Queries.GetCafes(ctx)
}
