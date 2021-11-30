package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponce struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

func initErrorResponce(c *gin.Context, statusCode int, errMessage string) {
	fmt.Printf("Json Error responce with message %s\n", errMessage)
	c.AbortWithStatusJSON(statusCode, ErrorResponce{Status: "Error", Message: errMessage})
}

func initOkResponce(c *gin.Context, params map[string]interface{}) {
	params["Status"] = "OK"

	c.JSON(http.StatusOK, params)
}
