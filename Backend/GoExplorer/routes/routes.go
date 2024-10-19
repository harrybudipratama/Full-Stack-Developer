package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/getFolders", getFolders)
	server.GET("/getFiles/:id", getFile)

}
