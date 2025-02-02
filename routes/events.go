package routes

import (
	"net/http"
	"strconv"

	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't fetch events please try again later",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "success",
		"events":  events,
	})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})

		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func createEvent(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data"})
		return
	}

	event.UserID = context.GetInt64("userId")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't store event please try again later",
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event was created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})

		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event wasn't found",
		})

		return
	}

	if event.UserID != context.GetInt64("userId") {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid ownership",
		})

		return
	}

	var updateEvent models.Event

	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data"})
		return
	}

	updateEvent.ID = eventId

	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Updating event has failed",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event Updated successfully",
		"event":   updateEvent,
	})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})

		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event wasn't found",
		})

		return
	}

	if event.UserID != context.GetInt64("userId") {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid ownership",
		})

		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Deleting event has failed",
		})

		return
	}

	context.JSON(http.StatusNoContent, nil)
}
