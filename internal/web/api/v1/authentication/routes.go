package authentication

import (
	"encoding/json"
	"github.com/MrDienns/gin-example/api/auth"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// RegisterRoutes will register the API routes for the /authentication relative path
func RegisterRoutes(routerGroup *gin.RouterGroup) {
	route := routerGroup.Group("/authenticate")

	route.POST("/", postAuthenticate)
}

func postAuthenticate(context *gin.Context) {

	bodyByteBuffer, err := ioutil.ReadAll(context.Request.Body)
	if len(bodyByteBuffer) == 0 {
		writeError(context, 400, "request body is empty")
		return
	}
	if err != nil {
		writeError(context, 400, err.Error())
		return
	}

	var authenticationResponse auth.AuthenticationResponse
	err = json.Unmarshal(bodyByteBuffer, &authenticationResponse)
	if err != nil {
		writeError(context, 400, err.Error())
		return
	}

	context.JSON(200, auth.AuthenticationResponse{
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
	})
}

func writeError(context *gin.Context, code int, err string) {
	context.JSON(code, gin.H{
		"error": err,
	})
}
