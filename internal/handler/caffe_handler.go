package handler

import (
	"net/http"

	"caffinity/config"
	"caffinity/internal/db"
	"caffinity/internal/models"
	"caffinity/internal/service"
	"caffinity/utils"

	"github.com/gin-gonic/gin"
)

type CafeHandler struct {
	service *service.CafeService
	config  *config.DBConfig
}

func NewCafeHandler(s *service.CafeService, cfg *config.DBConfig) *CafeHandler {
	return &CafeHandler{service: s, config: cfg}
}

func (h *CafeHandler) GetCafes(c *gin.Context) {
	var (
		data          []db.Cafe
		err           error
		city          string
		limit, offset int32
		message       = "success"
		response      = []models.CaffeResponse{}
	)

	limitString := c.Query("limit")
	offsetString := c.Query("offset")
	city = c.Query("city")

	limit = utils.ParseStringToInt32(limitString, h.config.LIMIT)
	offset = utils.ParseStringToInt32(offsetString, h.config.OFFSET)

	if city != "" {
		data, err = h.service.GetCafesByCity(c, city, limit, offset)
	} else {
		data, err = h.service.GetCafes(c)
	}

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
