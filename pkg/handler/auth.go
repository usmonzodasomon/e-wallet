package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/e-wallet/models"
)

func (h *handler) signUp(c *gin.Context) {
	var User models.User
	if err := c.BindJSON(&User); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Authorization.SignUp(&User); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var Wallet models.Wallet
	Wallet.UserID = User.ID
	if err := h.services.Wallet.CreateWallet(&Wallet); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      User.ID,
	})
}

func (h *handler) signIn(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
}
