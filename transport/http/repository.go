package http

import (
	"github.com/muhsatrio/golang-boilerplate/interactors/auth"
	"github.com/muhsatrio/golang-boilerplate/interactors/user"

	"github.com/gin-gonic/gin"
)

type HTTP struct {
	Config      Config
	Auth        Auth
	UserService user.Interactors
	AuthService auth.Interactors
}

func (h HTTP) Serve() {
	r := gin.Default()

	r.GET("/", healthCheck)

	v1 := r.Group("/api/v1")

	usersGroup := v1.Group("/users")
	usersGroup.POST("", h.userRegister)

	authGroup := v1.Group("/auth")
	authGroup.POST("/login", h.authLogin)

	r.Run(h.Config.Port)

}

func healthCheck(c *gin.Context) {
	c.JSON(200, "OK")
}
