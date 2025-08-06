package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/slimreaper35/konflux-test/models"
)

func GetOneEventHandler(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	event, err := models.GetEventBy(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func GetAllEventsHandler(context *gin.Context) {
	var events, err = models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get events"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func CreateEventHandler(context *gin.Context) {
	var userID = context.GetInt64("userID")
	var event models.Event = models.Event{UserID: userID}

	if context.ShouldBindJSON(&event) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if event.Create() != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create event"})
		return
	}

	context.JSON(http.StatusCreated, event)
}

func UpdateEventHandler(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	event, err := models.GetEventBy(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	var userID = context.GetInt64("userID")
	if event.UserID != userID {
		context.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	if context.ShouldBindJSON(&event) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if event.Update() != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func DeleteEventHandler(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	event, err := models.GetEventBy(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	var userID = context.GetInt64("userID")
	if event.UserID != userID {
		context.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	if event.Delete() != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete event"})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
