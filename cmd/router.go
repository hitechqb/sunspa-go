package cmd

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sunspa/common"
	"sunspa/module/att"
	"sunspa/sdk"
)

const (
	NAME    = "SunSpa App"
	VERSION = "1.0"
)

func Router(sc sdk.Service) {
	logger := sc.MustGet(common.KeyMainLogger).(common.Logger)
	logger.Infof("🎉 Running App [%s] - 🎉 version: %s", NAME, VERSION)

	engine := sc.EngineGin()
	engine.Use(CORSMiddleware())

	/// Version v1 for Endpoint from [A Phat]
	v1 := engine.Group("/api/v1")
	{
		v1.GET("/ping", Ping("pong"))

		// handle Auth
		user := v1.Group("/user/:id")
		{
			user.Use(AuthMiddleware())
			user.GET("/att", att.GetAttByUserId(sc))
		}
	}
}

func Ping(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, msg)
	}
}
