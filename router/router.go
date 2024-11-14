package router

import "github.com/gin-gonic/gin"

func  Initialize(){
	// Inicialita o Router com as configs Default do Gin
    router := gin.Default()
		//Definindo Rota
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H {
				"message" : "pong",
			})
		})
		// Rodando API
		router.Run(":8080")
}