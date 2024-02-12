package router

import (
	"receipt/internal/handler"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/receipts")
	{
		v1.POST("/process", handler.ReceiptsProcess)
		v1.GET("/:id/points", handler.ReceiptsPoints)
	}

	return r
}
