package web

import (
	"github.com/MrDienns/gin-example/internal/web/api"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type server struct {
	engine *gin.Engine
}

// NewServer will create a new *server.
func NewServer() *server {
	return &server{}
}

// Start will start a webserver. This method will block further code execution until the webserver is shut down.
func (s *server) Start() error {
	gin.SetMode(gin.ReleaseMode)
	s.engine = gin.Default()

	prom := ginprometheus.NewPrometheus("gin")
	prom.Use(s.engine)

	rootGroup := s.engine.Group("/")
	api.RegisterRoutes(rootGroup)

	return s.engine.Run()
}
