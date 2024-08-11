package main

import (
	"net/http"

	"example.com/events/db"
	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

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

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't store event please try again later",
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event was created", "event": event})
}
