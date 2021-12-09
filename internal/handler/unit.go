package handler

import (
	"encoding/json"
	"fmt"
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
	if err != nil {
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

func (h *Handler) intersectUnits(context *gin.Context) {
	type elements struct {
		First_id  int `json:"first_id"`
		Second_id int `json:"second_id"`
	}

	var sample elements
	err := context.BindJSON(&sample)
	if err != nil {
		fmt.Println("Error %v", err)
	}

	_, bytes, err := h.services.IntersectUnits(sample.First_id,sample.Second_id)
	if err != nil {
		initErrorResponce(context, http.StatusInternalServerError, err.Error())
		return
	}
	//context.JSON(http.StatusOK, res)
	context.JSON(http.StatusOK, json.RawMessage(*bytes))
}
