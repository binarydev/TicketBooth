package controllers

import (
	"github.com/binarydev/ticketbooth/datastore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllMediaRequests(ctx *gin.Context) {
	var reqs []datastore.MediaRequest
	db := datastore.OpenDatabaseConnection()
	db.Find(&datastore.MediaRequest{}, &reqs)
	for _, a := range reqs {
		ctx.IndentedJSON(http.StatusOK, a)
	}
}
