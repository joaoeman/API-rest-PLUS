package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOPeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}
