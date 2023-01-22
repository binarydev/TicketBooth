package controllers

import (
	"github.com/binarydev/ticketbooth/datastore"
	"github.com/binarydev/ticketbooth/models"
	"github.com/gin-gonic/gin"
	"log"
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

func GetRemoteIndex(ctx *gin.Context) {
	var model models.RemoteIndex
	id := ctx.Param("id")
	db := datastore.OpenDatabaseConnection()
	result := db.Find(&model, id)
	if result.Error != nil {
		log.Fatal(result.Error)
		ctx.JSON(http.StatusBadRequest, gin.H{})
	} else if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{})
	} else {
		ctx.JSON(http.StatusOK, model)
	}
}

func CreateRemoteIndex(ctx *gin.Context) {

}

func UpdateRemoteIndex(ctx *gin.Context) {
	
}
