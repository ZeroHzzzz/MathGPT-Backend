package utils

import (
	"mathgpt/app/utils/stateCode"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"code":    stateCode.OK,
		"message": "OK",
	})
}
