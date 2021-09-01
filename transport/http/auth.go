package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhsatrio/golang-boilerplate/interactors/auth"
)

func (h HTTP) authLogin(c *gin.Context) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		Token string `json:"token"`
	}

	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(errorHandler(err))
		return
	}

	reqService := auth.RequestLogin{
		Email:    req.Email,
		Password: req.Password,
		SaltHash: h.Auth.SaltHash,
	}

	result, err := h.AuthService.Login(reqService)
	if err != nil {
		c.AbortWithStatusJSON(errorHandler(err))
		return
	}

	resp := response{
		Token: result.Token,
	}

	c.JSON(http.StatusOK, resp)
}
