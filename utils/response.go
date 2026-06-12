package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
)

func SuccessResponse(
	c *gin.Context,
	statusCode int,
	message string,
	data interface{},
){
	c.JSON(statusCode,gin.H{
		"success": true,
		"message": message,
		"data":data,
	})
}

func ErrorResponse(
	c *gin.Context,
	statusCode int,
	message string,
){
	c.JSON(statusCode,gin.H{
		"success":false,
		"message":message,
	})
}


func InternalServerError(c *gin.Context) {

	ErrorResponse(
		c,
		http.StatusInternalServerError,
		"internal server error",
	)
}
