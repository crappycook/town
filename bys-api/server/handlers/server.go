package handlers

import (
	"bys/bootstrap"
	"bys/pkg/ulog"
	"bys/server/middleware"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	webEngine  *gin.Engine
	bootConfig *bootstrap.Config
}

type Option func(*HttpServer)

func NewHttpServer(options ...Option) *HttpServer {
	server := new(HttpServer)
	for _, f := range options {
		f(server)
	}
	return server
}

func WithBootConfig(cfg *bootstrap.Config) func(*HttpServer) {
	return func(s *HttpServer) {
		s.bootConfig = cfg
	}
}

func (s *HttpServer) InitRouter() {
	s.webEngine = gin.New()
	s.webEngine.Use(
		middleware.CROSHandler(),
		middleware.RecoveryHandler(),
	)

	s.webEngine.OPTIONS("/:path/", options)
	s.webEngine.OPTIONS("/:path/:path", options)
	s.webEngine.OPTIONS("/:path/:path/:path", options)
	s.webEngine.OPTIONS("/:path/:path/:path/:path", options)
	s.webEngine.OPTIONS("/:path/:path/:path/:path/:path", options)

	v1 := s.webEngine.Group("/api/v1")

	// common
	common := v1.Group("/common")
	common.GET("/ping", s.CommonConfig)
	common.GET("/host_status", s.GetRemoteHostStatus)

	// matches
	matches := v1.Group("/football_match")
	matches.GET("/:id", s.GetFootballMatchDetail)
}

func (s *HttpServer) Run() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	s.InitRouter()

	startMsg := fmt.Sprintf("run http server %v", s.bootConfig.RestApi.ListenPort)
	ulog.Info(startMsg)
	err := s.webEngine.Run(s.bootConfig.RestApi.ListenPort)
	if err != nil {
		panic(err)
	}
}

func options(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
