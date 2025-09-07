package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoeman/API-rest-PLUS/handler"
)

func inicializeRouters(router *gin.Engine) {
	v1 := router.Group("/api/v1") //adicionando um grupo, para indicar rotas

	v1.GET("/opening", handler.GetOPeningHandler)

	v1.POST("/opening", handler.PostOPeningHandler)

	v1.DELETE("/opening", handler.DeleteOPeningHandler)

	v1.PUT("/opening", handler.UpdateOPeningHandler)

	v1.GET("/openings", handler.GetsOPeningHandler)

}
