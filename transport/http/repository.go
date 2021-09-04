package http

import (
	"linkaja-test/interactors/account"

	"github.com/gin-gonic/gin"
)

type HTTP struct {
	Config             Config
	Auth               Auth
	AccountInteractors account.Interactors
}

func (h HTTP) Serve() {
	r := gin.Default()

	r.GET("/", healthCheck)

	accountGroup := r.Group("/account")

	accountGroup.GET("/:account_number", h.checkBalance)

	accountGroup.POST("/:from_account_number/transfer", h.transfer)

	r.Run(h.Config.Port)

}

func healthCheck(c *gin.Context) {
	c.JSON(200, "OK")
}
