package routes

import (
	"example/rest/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get events",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"error": err.Error()})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get event",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not create event",
			"error": err.Error()})
		return
	}
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not save event",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event created",
		"event": event,
	})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"error": err.Error()})
		return
	}
	_, err = models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get event",
			"error": err.Error()})
		return
	}
	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data",
			"error": err.Error()})
		return
	}
	event.ID = id
	err = event.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update event",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
		"event": event,
	})
}

func deleteEvent (c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
			"error": err.Error()})
		return
	}
	_, err = models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get event",
			"error": err.Error()})
		return
	}
	err = models.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not delete event",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted",
	})
}