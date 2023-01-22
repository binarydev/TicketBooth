package controllers

import (
	"github.com/binarydev/ticketbooth/datastore"
	"github.com/binarydev/ticketbooth/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllMediaRequests(ctx *gin.Context) {
	var reqs []models.MediaRequest
	db := datastore.OpenDatabaseConnection()
	db.Find(&models.MediaRequest{}, &reqs)
	for _, a := range reqs {
		ctx.JSON(http.StatusOK, a)
	}
}

func GetMediaRequest(ctx *gin.Context) {
	model := models.MediaRequest{}
	id := ctx.Param("id")
	db := datastore.OpenDatabaseConnection()
	db.Find(&model, id)
	ctx.JSON(http.StatusOK, model)
}

func CreateMediaRequest(ctx *gin.Context) {

}

func UpdateMediaRequest(ctx *gin.Context) {
	
}