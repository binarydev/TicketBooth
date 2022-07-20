package controllers

import (
	"github.com/binarydev/ticketbooth/datastore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllRemoteIndices(ctx *gin.Context) {
	var reqs []datastore.RemoteIndex
	db := datastore.OpenDatabaseConnection()
	db.Find(&reqs)
	for _, a := range reqs {
		ctx.IndentedJSON(http.StatusOK, a)
	}
}
