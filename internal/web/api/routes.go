package api

import (
	v1 "github.com/MrDienns/gin-example/internal/web/api/v1"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes will create a relative /api path, and registers different API version paths under it
func RegisterRoutes(routerGroup *gin.RouterGroup) {
	route := routerGroup.Group("/api")

	v1.RegisterRoutes(route)
}
