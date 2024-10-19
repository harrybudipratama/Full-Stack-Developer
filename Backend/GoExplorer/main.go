package main

import (
	"go-explorer/db"
	"go-explorer/routes"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT , DELETE , PATCH")
		c.Next()
	}
}

func main() {
	db.InitDB()
	server := gin.Default()
	server.Use(CORSMiddleware())
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
