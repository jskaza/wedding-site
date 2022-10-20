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

func RSVP(name string, info []map[string]string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, _ := mongo.Connect(context.TODO(), clientOptions)

	coll := client.Database("wedding").Collection("guests")
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "meals", Value: info}}}}
	result, _ := coll.UpdateOne(context.TODO(), models.Guest{DisplayName: name}, update)
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
}
