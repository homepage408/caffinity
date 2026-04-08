package handler

import (
	"net/http"

	"caffinity/internal/db"
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
		message = "success"
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

	c.JSON(http.StatusOK, utils.NewResponse(http.StatusOK, message, data))
}
