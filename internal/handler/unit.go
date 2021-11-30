package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DimKush/geometry_go/internal/entity/unit"
)

func (h *Handler) setUnit(context *gin.Context) {
	var unit unit.Unit

	if err := context.BindJSON(&unit); err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.SetUnit(unit)
	if err != nil {

	}
}