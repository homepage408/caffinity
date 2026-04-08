package repository

import (
	"caffinity/internal/db"
	"context"
	"strings"
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

func (r *CafeRepository) GetByCity(ctx context.Context, city string, limit int32, offset int32) ([]db.Cafe, error) {
	params := db.GetCaffesByCityParams{
		City:   strings.ToLower(city),
		Limit:  limit,
		Offset: offset,
	}
	return r.Queries.GetCaffesByCity(ctx, params)
}
