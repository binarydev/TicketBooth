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
