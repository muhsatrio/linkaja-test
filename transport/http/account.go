package http

import (
	"linkaja-test/interactors/account"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HTTP) checkBalance(c *gin.Context) {
	type response struct {
		AccountNumber uint   `json:"account_number"`
		CustomerName  string `json:"customer_name"`
		Balance       int    `json:"balance"`
	}

	accountNumber := c.Param("account_number")

	convertedAccNumber, _ := strconv.ParseUint(accountNumber, 10, 32)

	result, err := h.AccountInteractors.CheckBalance(uint(convertedAccNumber))

	if err != nil {
		c.AbortWithStatusJSON(errorHandler(err))
		return
	}

	resp := response{
		AccountNumber: result.AccountNumber,
		CustomerName:  result.CustomerName,
		Balance:       result.Balance,
	}

	c.JSON(http.StatusOK, resp)
}

func (h HTTP) transfer(c *gin.Context) {
	type request struct {
		ToAccountNumber uint `json:"to_account_number"`
		Amount          int  `json:"amount"`
	}

	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(errorHandler(err))
		return
	}

	senderAccountNumber := c.Param("from_account_number")

	convertedSenderAccNumber, _ := strconv.ParseUint(senderAccountNumber, 10, 32)

	reqService := account.RequestTransfer{
		SenderAccountNumber:   uint(convertedSenderAccNumber),
		ReceiverAccountNumber: req.ToAccountNumber,
		Amount:                req.Amount,
	}

	err := h.AccountInteractors.Transfer(reqService)
	if err != nil {
		c.AbortWithStatusJSON(errorHandler(err))
		return
	}

	c.Status(http.StatusNoContent)
}
