package service

import (
	"caffinity/internal/models"
	"caffinity/internal/repository"
	"context"
)

type CafeService interface {
	// Cafes
	GetCafes(ctx context.Context, limit, offset int32) ([]models.CafeResponse, error)
	GetCafesByCity(ctx context.Context, city string, limit int32, offset int32) ([]models.CafeResponse, error)
}
type cafeService struct {
	repo repository.CafeRepository
}

func NewCafeService(r repository.CafeRepository) *cafeService {
	return &cafeService{repo: r}
}

func (s *cafeService) GetCafes(ctx context.Context, limit, offset int32) ([]models.CafeResponse, error) {
	var (
		ids          []int32
		cafeResponse []models.CafeResponse
	)

	datas, err := s.repo.GetAllCafes(ctx, limit, offset)
	if err != nil {
		return []models.CafeResponse{}, err
	}

	for _, data := range datas {
		ids = append(ids, data.ID)
	}

	facilities, err := s.repo.GetFacilitiesByCafeIDs(ctx, ids)
	if err != nil {
		return []models.CafeResponse{}, err
	}

	tags, err := s.repo.GetTagsByCafeIDs(ctx, ids)
	if err != nil {
		return []models.CafeResponse{}, err
	}

	menus, err := s.repo.GetMenusByCafeIDs(ctx, ids)
	if err != nil {
		return []models.CafeResponse{}, err
	}

	facilitiesMap := make(map[int32][]models.Facilities)
	for _, facility := range facilities {
		facilitiesMap[facility.CafeID] = append(facilitiesMap[facility.CafeID], models.Facilities{
			ID:     facility.ID,
			CafeID: facility.CafeID,
			Name:   facility.Name,
		})
	}

	tagsMap := make(map[int32][]models.Tags)
	for _, tag := range tags {
		tagsMap[tag.CafeID] = append(tagsMap[tag.CafeID], models.Tags{
			ID:     tag.TagID,
			CafeID: tag.CafeID,
			Name:   tag.Name,
		})
	}

	menusMap := make(map[int32][]models.Menus)
	for _, menu := range menus {
		menusMap[menu.CafeID] = append(menusMap[menu.CafeID], models.Menus{
			ID:          menu.ID,
			CafeID:      menu.CafeID,
			Name:        menu.Name,
			Price:       menu.Price,
			Strength:    menu.Strength.Int32,
			IsSafe:      menu.IsSafe.Bool,
			Description: menu.Description.String,
			Image:       menu.Image.String,
		})
	}

	for _, data := range datas {
		cafeResponse = append(cafeResponse, models.CafeResponse{
			ID:             data.ID,
			Name:           data.Name,
			TagLine:        data.Tagline.String,
			Address:        data.Address.String,
			Lat:            data.Latitude.String,
			Lng:            data.Longitude.String,
			Hours:          data.OpenHours.String,
			Phone:          data.Phone.String,
			Instagram:      data.Instagram.String,
			Rating:         data.Rating.String,
			Reviews:        data.Reviews.Int32,
			PriceLevel:     data.PriceLevel.Int32,
			Approved:       data.Approved.Bool,
			ApprovedReason: data.ApprovedReason.String,
			HeroImage:      data.HeroImg.String,

			Facilities: facilitiesMap[data.ID],
			Tags:       tagsMap[data.ID],
			Menus:      menusMap[data.ID],
		})
	}

	return cafeResponse, nil
}

func (s *cafeService) GetCafesByCity(ctx context.Context, city string, limit int32, offset int32) ([]models.CafeResponse, error) {

	var (
		ids          []int32
		cafeResponse []models.CafeResponse
	)

	datas, err := s.repo.GetCafesByCity(ctx, city, limit, offset)
	if err != nil {
		return []models.CafeResponse{}, err
	}

	for _, data := range datas {
		ids = append(ids, data.ID)
	}

	facilities, err := s.repo.GetFacilitiesByCafeIDs(ctx, ids)
	if err != nil {
		return []models.CafeResponse{}, err
	}

	tags, err := s.repo.GetTagsByCafeIDs(ctx, ids)
	if err != nil {
		return []models.CafeResponse{}, err
	}

	menus, err := s.repo.GetMenusByCafeIDs(ctx, ids)
	if err != nil {
		return []models.CafeResponse{}, err
	}

	facilitiesMap := make(map[int32][]models.Facilities)
	for _, facility := range facilities {
		facilitiesMap[facility.CafeID] = append(facilitiesMap[facility.CafeID], models.Facilities{
			ID:     facility.ID,
			CafeID: facility.CafeID,
			Name:   facility.Name,
		})
	}

	tagsMap := make(map[int32][]models.Tags)
	for _, tag := range tags {
		tagsMap[tag.CafeID] = append(tagsMap[tag.CafeID], models.Tags{
			ID:     tag.TagID,
			CafeID: tag.CafeID,
			Name:   tag.Name,
		})
	}

	menusMap := make(map[int32][]models.Menus)
	for _, menu := range menus {
		menusMap[menu.CafeID] = append(menusMap[menu.CafeID], models.Menus{
			ID:          menu.ID,
			CafeID:      menu.CafeID,
			Name:        menu.Name,
			Price:       menu.Price,
			Strength:    menu.Strength.Int32,
			IsSafe:      menu.IsSafe.Bool,
			Description: menu.Description.String,
			Image:       menu.Image.String,
		})
	}

	for _, data := range datas {
		cafeResponse = append(cafeResponse, models.CafeResponse{
			ID:             data.ID,
			Name:           data.Name,
			TagLine:        data.Tagline.String,
			Address:        data.Address.String,
			Lat:            data.Latitude.String,
			Lng:            data.Longitude.String,
			Hours:          data.OpenHours.String,
			Phone:          data.Phone.String,
			Instagram:      data.Instagram.String,
			Rating:         data.Rating.String,
			Reviews:        data.Reviews.Int32,
			PriceLevel:     data.PriceLevel.Int32,
			Approved:       data.Approved.Bool,
			ApprovedReason: data.ApprovedReason.String,
			HeroImage:      data.HeroImg.String,

			Facilities: facilitiesMap[data.ID],
			Tags:       tagsMap[data.ID],
			Menus:      menusMap[data.ID],
		})
	}

	return cafeResponse, nil
}
