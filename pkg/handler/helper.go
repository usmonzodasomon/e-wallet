package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewErrorResponse(c *gin.Context, status int, err string) {
	c.AbortWithStatusJSON(status, errorResponse{
		Message: "error",
		Error:   err,
	})
}

func GetUserId(c *gin.Context) (int64, error) {
	userId, ok := c.Get("X-UserId")
	if !ok {
		return 0, errors.New("error while getting userId from header")
	}
	id, ok := userId.(int64)
	if !ok {
		return 0, errors.New("error while convertation userId to int64")
	}
	return id, nil
}
