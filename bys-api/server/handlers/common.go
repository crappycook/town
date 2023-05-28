package handlers

import (
	"bys/pkg/ulog"
	"bys/rpc"
	"bys/util/resp"

	"github.com/gin-gonic/gin"
)

func (s *HttpServer) CommonConfig(gCtx *gin.Context) {
	ulog.Info("api recv ping req")
	resp.SuccessWithData(gCtx, "ping-pong")
}

func (s *HttpServer) GetRemoteHostStatus(gCtx *gin.Context) {
	ulog.Info("recv get remote host status req")
	rsp, err := rpc.GetHostStatus(gCtx)
	if err != nil {
		ulog.Error("rpc.GetHostStatus error", ulog.String("error", err.Error()))
		return
	}
	resp.SuccessWithData(gCtx, map[string]interface{}{
		"host_name": rsp.GetHostname(),
		"time_ms":   rsp.GetTimestampMs(),
	})
}
