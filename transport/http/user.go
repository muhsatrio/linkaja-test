package http

import (
	"net/http"

	"github.com/muhsatrio/golang-boilerplate/interactors/user"

	"github.com/gin-gonic/gin"
)

func (h HTTP) userRegister(c *gin.Context) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	type response struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(errorHandler(err))
		return
	}

	reqService := user.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		SaltHash: h.Password.SaltHash,
	}

	result, err := h.UserService.Register(reqService)
	if err != nil {
		c.AbortWithStatusJSON(errorHandler(err))
		return
	}

	resp := response{
		ID:    result.ID,
		Email: result.Email,
		Name:  result.Name,
	}

	c.JSON(http.StatusOK, resp)
}
