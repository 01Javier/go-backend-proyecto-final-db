package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Status:  "error",
		Message: message,
	})
}

func ErrorResponseWithDetail(c *gin.Context, message string, statusCode int, err error) {
	response := Response{
		Status:  "error",
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(statusCode, response)
}
