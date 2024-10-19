package routes

import (
	"go-explorer/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getFolders(context *gin.Context) {
	folders, err := models.GetAllFolders()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch folders. Try again later."})
		return
	}
	context.JSON(http.StatusOK, folders)
}

func getFile(context *gin.Context) {
	fileId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse file id"})
		return
	}

	files, err := models.GetFileById(fileId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch files"})
		return
	}
	context.JSON(http.StatusOK, files)

}
