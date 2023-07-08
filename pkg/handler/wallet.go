package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
