package routes

import (
	"log"
	"net/http"
	"strconv"

	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
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

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register user",
		})

		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})

		return
	}

	event := models.Event{
		ID: eventId,
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't cancel registration",
		})

		return
	}

	context.JSON(http.StatusNoContent, nil)

}
