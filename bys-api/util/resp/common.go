package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessWithData(gCtx *gin.Context, data interface{}) {
	result := &JsonResult{
		Code:    Success,
		Message: "success",
		Data:    data,
	}
	gCtx.JSON(http.StatusOK, result)
}

func JsonOutputError(gCtx *gin.Context, ec int, em string, data interface{}) {
	result := &JsonResult{
		Code:    ec,
		Message: em,
		Data:    data,
	}
	gCtx.JSON(http.StatusOK, result)
}
