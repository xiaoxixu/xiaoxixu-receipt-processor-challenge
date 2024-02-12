package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonError(c *gin.Context, obj any) {
	c.JSON(http.StatusBadRequest, obj)
}

func JsonNotFoundError(c *gin.Context, obj any) {
	c.JSON(http.StatusNotFound, obj)
}

func JsonOk(c *gin.Context, obj any) {
	c.JSON(http.StatusOK, obj)
}
