package repository

import (
	"caffinity/internal/db"
	"context"
	"database/sql"
)

type CafeRepository interface {

	// Cafes
	GetAllCafes(ctx context.Context, limit int32, offset int32) ([]db.Cafe, error)
	GetCafesByCity(ctx context.Context, city string, limit int32, offset int32) ([]db.Cafe, error)

	// Facilities
	GetFacilitiesByCafeIDs(ctx context.Context, cafeIDs []int32) ([]db.GetFacilitiesByCafeIDsRow, error)

	// Tags
	GetTagsByCafeIDs(ctx context.Context, cafeIDs []int32) ([]db.GetTagsByCafeIDsRow, error)

	// Menus
	GetMenusByCafeIDs(ctx context.Context, cafeIDs []int32) ([]db.GetMenusByCafeIDsRow, error)
}

type cafeRepository struct {
	Queries *db.Queries
}

func NewCafeRepository(q *db.Queries) *cafeRepository {
	return &cafeRepository{Queries: q}
}

func (r *cafeRepository) GetAllCafes(ctx context.Context, limit int32, offset int32) ([]db.Cafe, error) {
	return r.Queries.GetAllCafes(ctx, db.GetAllCafesParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *cafeRepository) GetCafesByCity(ctx context.Context, city string, limit int32, offset int32) ([]db.Cafe, error) {

	cityParam := sql.NullString{String: city, Valid: city != ""}

	params := db.GetCafesByCityParams{
		Limit:  limit,
		Offset: offset,
		City:   cityParam,
	}

	return r.Queries.GetCafesByCity(ctx, params)
}

func (r *cafeRepository) GetFacilitiesByCafeIDs(ctx context.Context, cafeIDs []int32) ([]db.GetFacilitiesByCafeIDsRow, error) {
	return r.Queries.GetFacilitiesByCafeIDs(ctx, cafeIDs)
}

func (r *cafeRepository) GetTagsByCafeIDs(ctx context.Context, cafeIDs []int32) ([]db.GetTagsByCafeIDsRow, error) {
	return r.Queries.GetTagsByCafeIDs(ctx, cafeIDs)
}

func (r *cafeRepository) GetMenusByCafeIDs(ctx context.Context, cafeIDs []int32) ([]db.GetMenusByCafeIDsRow, error) {
	return r.Queries.GetMenusByCafeIDs(ctx, cafeIDs)
}
