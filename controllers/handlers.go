package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jskaza/wedding-site/helpers"
	"github.com/jskaza/wedding-site/models"
)

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func StoryGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "story.html", gin.H{})
	}
}

func PartyGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "party.html", gin.H{})
	}
}

func ScheduleGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "schedule.html", gin.H{})
	}
}

func RSVPGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		guests, _ := helpers.RetrieveGuests()
		c.HTML(http.StatusOK, "rsvp.html", guests)
	}
}

func RSVPPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json models.RSVPData
		c.BindJSON(&json)
		helpers.RSVP(json)
	}
}
