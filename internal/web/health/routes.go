package health

import "github.com/gin-gonic/gin"

// RegisterRoutes will register a Kubernetes compatible /health endpoint
func RegisterRoutes(routerGroup *gin.RouterGroup) {
	route := routerGroup.Group("/health")

	route.GET("", func(context *gin.Context) {

		full := context.DefaultQuery("full", "0")
		checks := gin.H{}

		if full != "0" {
			checks = gin.H{
				"webserver": "OK",
			}
		}
		context.JSON(200, checks)
	})
}
