package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jskaza/wedding-site/helpers"
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
		c.HTML(http.StatusOK, "schedule.html", gin.H{})
	}
}

func RSVPPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := make([]map[string]string, 0)
		data = append(data, map[string]string{"Jon": "Chicken"})
		helpers.RSVP("Mr. 88.117", data)
		c.MultipartForm()
		for key, value := range c.Request.PostForm {
			log.Printf("%v = %v \n", key, value)
		}
	}
}

// func MealDist() {
// 	clientOptions := options.Client().ApplyURI(uri)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	coll := client.Database("wedding").Collection("guests")

// 	cursor, err := coll.Find(context.TODO(), models.Guest{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var guests []models.Guest
// 	if err = cursor.All(context.TODO(), &guests); err != nil {
// 		log.Fatal(err)
// 	}
// 	meals := make([]string, 0)
// 	for i := range guests {
// 		for _, j := range guests[i].Meals {
// 			meals = append(meals, j["meal"])
// 		}
// 	}
// 	fmt.Println(meals)
// }

// func CheckPW(pw string) bool {
// 	clientOptions := options.Client().ApplyURI(uri)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	coll := client.Database("wedding").Collection("settings")
// 	var res models.Settings
// 	err = coll.FindOne(context.TODO(), models.Settings{}).Decode(&res)

// 	return pw == res.Password
// }
