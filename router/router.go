package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default() //inicializa o router usando as configs default do gin

	inicializeRouters(router)
	//rodando servidor
	router.Run(":8080")
}
