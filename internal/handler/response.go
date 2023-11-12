package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json: "status"`
}

func newErrorResponse(c *gin.Context, StatusCode int, message string) {
	log.Print(message)
	c.AbortWithStatusJSON(StatusCode, errorResponse{message})
}
