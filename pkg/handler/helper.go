package handler

import "github.com/gin-gonic/gin"

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
