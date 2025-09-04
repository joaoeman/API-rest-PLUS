package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()                    //inicializa o router usando as configs default do gin
	router.GET("/ping", func(c *gin.Context) { //contexto da api
		c.JSON(200, gin.H{ //é o encode, transformando oq recebe em json
			"message": "pong",
		})
	})
	router.Run(":8080")
}
