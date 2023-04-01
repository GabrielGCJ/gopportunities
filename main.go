package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Criando o objeto gin.Engine para rotas HTTP
	r := gin.Default()
	// Definindo uma rota com o método GET
	r.GET("/ping", func(c *gin.Context) {
		// Definindo a resposta JSON para a rota
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Iniciando o servidor HTTP
	r.Run()
}