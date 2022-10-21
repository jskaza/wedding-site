package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jskaza/wedding-site/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = os.Getenv("MONGO_WEDDING")

func RetrieveGuests() ([]models.Guest, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	coll := client.Database("wedding").Collection("guests")

	cursor, err := coll.Find(context.TODO(), models.Guest{})
	if err != nil {
		log.Fatal(err)
	}
	var guests []models.Guest
	if err = cursor.All(context.TODO(), &guests); err != nil {
		log.Fatal(err)
	}

	return guests, nil
}

func RSVP(data models.RSVPData) {
	clientOptions := options.Client().ApplyURI(uri)
	client, _ := mongo.Connect(context.TODO(), clientOptions)

	meals := make([]map[string]string, 0)
	for _, v := range data.People {
		meals = append(meals, v)
	}

	coll := client.Database("wedding").Collection("guests")
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "meals", Value: meals}}}}
	result, _ := coll.UpdateOne(context.TODO(), models.Guest{DisplayName: data.DisplayName}, update)
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
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
