package controllers

import (
	"context"
	"log"
	"os"

	"github.com/jskaza/wedding-site/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = os.Getenv("MONGO_WEDDING")

func RetrieveGuestByName(name string) (string, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	query := bson.M{
		"$or": []interface{}{
			bson.M{"guests": name},
			bson.M{"displayName": name},
		},
	}
	coll := client.Database("wedding").Collection("guests")
	var res models.Guest
	err = coll.FindOne(context.TODO(), query).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", err
		}
		panic(err)
	}
	return res.ID.Hex(), nil
}

func RetrieveGuestByID(id string) (models.Guest, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("wedding").Collection("guests")
	var res models.Guest
	err = coll.FindOne(context.TODO(), models.Guest{ID: objectId}).Decode(&res)

	return res, nil
}

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

func CheckPW(pw string) bool {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("wedding").Collection("settings")
	var res models.Settings
	err = coll.FindOne(context.TODO(), models.Settings{}).Decode(&res)

	return pw == res.Password
}
