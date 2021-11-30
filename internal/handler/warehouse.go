package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getWarehouseById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	fmt.Println(id)
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}
	whs, err := h.services.Warehouse.GetWarehouseById(id)
	if err != nil {
		initErrorResponce(context, http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, whs)
}