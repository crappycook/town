package handlers

import (
	"bys/dal"
	"bys/util/resp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *HttpServer) GetFootballMatchDetail(gCtx *gin.Context) {
	var id int64
	idStr := gCtx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		resp.JsonOutputError(gCtx, resp.InvalidParams, err.Error(), nil)
		return
	}
	data, err := dal.GetFootballMatchByID(dal.LocalDB, id)
	if err != nil {
		resp.JsonOutputError(gCtx, resp.DBError, err.Error(), nil)
		return
	}
	resp.SuccessWithData(gCtx, data)
}
