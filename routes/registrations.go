package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/slimreaper35/konflux-test/models"
)

func RegisterForEventHandler(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	event, err := models.GetEventBy(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	var userID = context.GetInt64("userID")

	if event.RegisterUser(userID) != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

func UnregisterFromEventHandler(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	event, err := models.GetEventBy(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	var userID = context.GetInt64("userID")

	if event.UnregisterUser(userID) != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unregister from event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "OK"})
}
