package controllers

import (
	"github.com/binarydev/ticketbooth/datastore"
	"github.com/binarydev/ticketbooth/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllRemoteIndices(ctx *gin.Context) {
	var indices []models.RemoteIndex
	db := datastore.OpenDatabaseConnection()
	db.Find(&indices)
	for _, a := range indices {
		ctx.JSON(http.StatusOK, a)
	}
}
