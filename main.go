package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jskaza/wedding-site/controllers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./views/*.html")
	router.Static("./assets/css", "./assets/css")
	router.Static("./assets/img", "./assets/img")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/story", func(c *gin.Context) {
		c.HTML(http.StatusOK, "story.html", gin.H{})
	})

	router.GET("/party", func(c *gin.Context) {
		c.HTML(http.StatusOK, "party.html", gin.H{})
	})

	router.GET("/schedule", func(c *gin.Context) {
		c.HTML(http.StatusOK, "schedule.html", gin.H{})
	})

	router.GET("/rsvp", func(c *gin.Context) {
		guests, _ := controllers.RetrieveGuests()
		c.HTML(http.StatusOK, "rsvp.html", guests)
	})

	router.POST("/rsvp/lookup", func(c *gin.Context) {
		id, err := controllers.RetrieveGuestByName(c.PostForm("name"))
		if err != nil {
			c.String(http.StatusOK, "try again", gin.H{})
		} else {
			c.Redirect(http.StatusFound, "/rsvp/"+id)
		}
	})

	router.GET("/rsvp/:id", func(c *gin.Context) {
		guest, _ := controllers.RetrieveGuestByID(c.Param("id"))
		c.HTML(http.StatusOK, "rsvp_single.html", guest)
	})

	// probably put
	router.POST("/rsvp/guests", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/rsvp/"+c.PostForm("id")+"/meals")
	})

	router.GET("/rsvp/:id/meals", func(c *gin.Context) {
		guest, _ := controllers.RetrieveGuestByID(c.Param("id"))
		c.HTML(http.StatusFound, "rsvp_meals.html", guest)
	})

	// router.GET("/analytics", func(c *gin.Context) {
	// 	controllers.MealDist()
	// 	c.HTML(http.StatusFound, "analytics.html", gin.H{})
	// })

	router.POST("/rsvp/verify", func(c *gin.Context) {
		check := controllers.CheckPW(c.PostForm("pw"))
		if check {
			guests, _ := controllers.RetrieveGuests()
			c.HTML(http.StatusFound, "rsvp_single.html", guests)
		} else {
			fmt.Println("Incorrect PW")
		}

	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
