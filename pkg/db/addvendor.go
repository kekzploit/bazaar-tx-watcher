package db

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func AddVendor(mongoUri, image string, title string, description string, secret string, accType string, amount int64, txHash string) {
	// instantiate mongodb client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("marketplace").Collection("vendors")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.D{
		{Key: "image", Value: image},
		{Key: "title", Value: title},
		{Key: "description", Value: description},
		{Key: "socials", Value: bson.D{
			{Key: "twitter", Value: ""},
			{Key: "discord", Value: ""},
			{Key: "telegram", Value: ""},
			{Key: "github", Value: ""},
			{Key: "email", Value: ""},
		}},
		{Key: "offers", Value: bson.A{}},
		{Key: "secret", Value: secret},
		{Key: "url", Value: uuid.New().String()},
		{Key: "time", Value: time.Now().Unix()},
		{Key: "type", Value: accType},
		{Key: "amount", Value: amount},
		{Key: "tx_hash", Value: txHash}})

	id := res.InsertedID

	if id != "" {
		fmt.Println(id)
	} else {
		fmt.Println("no ObjectId present")
	}
}
