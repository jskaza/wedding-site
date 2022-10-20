package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jskaza/wedding-site/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.IndexGetHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/story", controllers.StoryGetHandler())
	g.GET("/party", controllers.PartyGetHandler())
	g.GET("/schedule", controllers.ScheduleGetHandler())
	g.GET("/rsvp", controllers.RSVPGetHandler())
	g.POST("/rsvp/guests", controllers.RSVPPostHandler())
}
