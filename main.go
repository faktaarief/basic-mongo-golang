package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type post struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("YOUR_MONGO_URI")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("basic-mongo"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("posts").InsertOne(ctx, post{"This A Title From Go", "Description test"})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print("Insert Success!")
}

func findAll() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	cursor, err := db.Collection("posts").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cursor.Close(ctx)

	result := make([]post, 0)

	for cursor.Next(ctx) {
		var row post
		err := cursor.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	fmt.Println("====================")

	for _, res := range result {
		fmt.Println("Judul:", res.Title)
		fmt.Println("Deskripsi:", res.Description)
		fmt.Println("====================")
	}
}

func find() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	objectId, err := primitive.ObjectIDFromHex("6114cd45ab8c6ef0bb489870")
	if err != nil {
		log.Fatal(err.Error())
	}

	cursor, err := db.Collection("posts").Find(ctx, bson.M{"_id": objectId})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cursor.Close(ctx)

	result := make([]post, 0)

	for cursor.Next(ctx) {
		var row post
		err := cursor.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	fmt.Println("====================")

	for _, res := range result {
		fmt.Println("Judul:", res.Title)
		fmt.Println("Deskripsi:", res.Description)
	}
}

func update() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	objectId, err := primitive.ObjectIDFromHex("6114cd45ab8c6ef0bb489870")
	if err != nil {
		log.Fatal(err.Error())
	}

	selector := bson.M{"_id": objectId}
	changes := post{"Update A Title", "Description Updated"}

	_, err = db.Collection("posts").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update success!")
}

func remove() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	objectId, err := primitive.ObjectIDFromHex("6114cd45ab8c6ef0bb489870")
	if err != nil {
		log.Fatal(err.Error())
	}

	selector := bson.M{"_id": objectId}
	result, err := db.Collection("posts").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove Success")
	fmt.Println(*result)
}

func main() {
	// insert()
	// findAll()
	// find()
	// update()
	// remove()
}
