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
	GetFacilitiesByCafeID(ctx context.Context, cafeID int32) ([]db.Facility, error)

	// Tags
	GetTagsByCafeID(ctx context.Context, cafeID int32) ([]db.GetTagsByCafeIDRow, error)

	// Menus
	GetMenusByCafeID(ctx context.Context, cafeID int32) ([]db.GetMenusByCafeIDRow, error)
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

func (r *cafeRepository) GetFacilitiesByCafeID(ctx context.Context, cafeID int32) ([]db.Facility, error) {
	return r.Queries.GetFacilitiesByCafeID(ctx, cafeID)
}

func (r *cafeRepository) GetTagsByCafeID(ctx context.Context, cafeID int32) ([]db.GetTagsByCafeIDRow, error) {
	return r.Queries.GetTagsByCafeID(ctx, cafeID)
}

func (r *cafeRepository) GetMenusByCafeID(ctx context.Context, cafeID int32) ([]db.GetMenusByCafeIDRow, error) {
	return r.Queries.GetMenusByCafeID(ctx, cafeID)
}
