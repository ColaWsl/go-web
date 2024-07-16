package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondJSON(c *gin.Context, status int, payload interface{}) {
	msg := "success"
	if status == http.StatusInternalServerError {
		msg = "error"
	}
	c.JSON(status, gin.H{
		"data": payload,
		"msg":  msg,
	})
}
