package web

import (
	"github.com/MrDienns/gin-example/internal/web/api"
	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
}

// NewServer will create a new *server.
func NewServer() *server {
	return &server{}
}

// Start will start a webserver. This method will block further code execution until the webserver is shut down.
func (s *server) Start() error {
	gin.SetMode(gin.ReleaseMode)
	s.router = gin.Default()

	api.RegisterRoutes(s.router.Group("/"))

	return s.router.Run()
}
