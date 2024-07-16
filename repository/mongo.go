package repository

import (
	"context"
	"fmt"
	"getAllUuid/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, context.Context, context.CancelFunc) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx, cancel
}

func GetAllUUIDs(client *mongo.Client, ctx context.Context) ([]model.UUID, error) {
	collection := client.Database("Uuid").Collection("getAllUuid")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	fmt.Println("Successfully connected and pinged.")

	var uuids []model.UUID
	for cur.Next(ctx) {
		var result model.UUID
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		uuids = append(uuids, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	// fmt.Println("Successfully connected and pinged.")

	return uuids, nil
}
