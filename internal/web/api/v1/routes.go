package v1

import (
	"github.com/MrDienns/gin-example/internal/web/api/v1/authentication"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes will create a relative /v1 group and register further underlying API routes
func RegisterRoutes(routerGroup *gin.RouterGroup) {
	route := routerGroup.Group("/v1")

	authentication.RegisterRoutes(route)
}
