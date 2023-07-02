package response

import (
	"github.com/gin-gonic/gin"
)

type Success struct {
	Data interface{}
}
type Error struct {
	Message string
	Details interface{}
	Code    int
}

func SendErrorResponse(s Error) gin.H {
	return gin.H{
		"status":  "error",
		"message": s.Message,
		"code":    s.Code,
		"details": s.Details,
	}
}

func SendSuccessResponse(s Success) gin.H {
	return gin.H{
		"status": "success",
		"data":   s.Data,
	}
}
