package handler

import (
	"net/http"

	"caffinity/internal/db"
	"caffinity/internal/models"
	"caffinity/internal/service"
	"caffinity/utils"

	"github.com/gin-gonic/gin"
)

type CafeHandler struct {
	service *service.CafeService
}

func NewCafeHandler(s *service.CafeService) *CafeHandler {
	return &CafeHandler{service: s}
}

func (h *CafeHandler) GetCafes(c *gin.Context) {
	var (
		message  = "success"
		response = []models.CaffeResponse{}
	)

	data, err := h.service.GetCafes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	if len(data) == 0 {
		message = "no cafes found"
		data = []db.Cafe{}
	}

	for _, v := range data {
		response = append(response, models.CaffeResponse{
			ID:          v.ID,
			Name:        v.Name,
			City:        v.City,
			Address:     utils.NullStringToPtr(v.Address),
			Description: utils.NullStringToPtr(v.Description),
			Rating:      utils.NullFloat64ToPtr(v.Rating),
			CreatedAt:   utils.NullTimeToPtr(v.CreatedAt),
		})
	}

	c.JSON(http.StatusOK, utils.NewResponse(http.StatusOK, message, response))
}
