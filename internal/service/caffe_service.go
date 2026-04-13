package service

import (
	"caffinity/internal/db"
	"caffinity/internal/repository"
	"context"
)

type CafeService struct {
	repo *repository.CafeRepository
}

func NewCafeService(r *repository.CafeRepository) *CafeService {
	return &CafeService{repo: r}
}

func (s *CafeService) GetCafes(ctx context.Context) ([]db.Cafe, error) {
	datas, err := s.repo.GetAll(ctx)
	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (s *CafeService) GetCafesByCity(ctx context.Context, city string, limit int32, offset int32) ([]db.Cafe, error) {
	datas, err := s.repo.GetByCity(ctx, city, limit, offset)
	if err != nil {
		return datas, err
	}

	return datas, nil
}
