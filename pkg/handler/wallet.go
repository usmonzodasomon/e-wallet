package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/e-wallet/models"
)

func (h *handler) checkAccount(c *gin.Context) {
	userID, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	Wallet, err := h.services.Wallet.CheckAccount(userID)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"wallet":  Wallet,
	})
}

func (h *handler) topUp(c *gin.Context) {
	userID, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var TopUp models.TopUp
	if err := c.BindJSON(&TopUp); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	TopUp.SenderID = userID

	transaction, err := h.services.Wallet.TopUp(TopUp)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"transaction": transaction,
	})
}

func (h *handler) getBalance(c *gin.Context) {
	userID, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	balance, err := h.services.Wallet.GetBalance(userID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"balance": balance,
	})
}

func (h *handler) monthStatistic(c *gin.Context) {
	userID, err := GetUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	totalCount, totalAmount, err := h.services.Wallet.MonthStatistic(userID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"total_count":  totalCount,
		"total_amount": totalAmount,
	})
}
