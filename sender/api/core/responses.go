package core

import (
	"github.com/gin-gonic/gin"
)

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Status     string    `json:"status"`
	StatusCode int       `json:"-"`
	Info       ErrorInfo `json:"info"`
}

func ErrorResponse(c *gin.Context, err Error) {
	c.IndentedJSON(err.StatusCode, err)
}

type Success struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func SuccessResponse(c *gin.Context, statusCode int, data any) {
	success := Success{}
	success.Status = "success"
	success.Data = data
	c.IndentedJSON(statusCode, success)
}
