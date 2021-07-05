package http

import (
	"financial-planner-be/interactors/user"

	"github.com/gin-gonic/gin"
)

type HTTP struct {
	Config      Config
	Password    Password
	UserService user.Service
}

func (h HTTP) Serve() {
	r := gin.Default()

	r.GET("/", healthCheck)

	v1 := r.Group("/api/v1")

	usersGroup := v1.Group("/users")
	usersGroup.POST("", h.userRegister)

	r.Run(h.Config.Port)

}

func healthCheck(c *gin.Context) {
	c.JSON(200, "OK")
}
