package http

import (
	"github.com/gin-gonic/gin"
)

type HTTP struct {
	Config Config
	Auth   Auth
}

func (h HTTP) Serve() {
	r := gin.Default()

	r.GET("/", healthCheck)

	r.Run(h.Config.Port)

}

func healthCheck(c *gin.Context) {
	c.JSON(200, "OK")
}
