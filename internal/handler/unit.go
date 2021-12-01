package handler

import (
	"net/http"
	"strconv"

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
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) getUnit(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err !=nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	unit, err := h.services.GetUnit(id)
	if err != nil || unit == nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, unit)
}