package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Socials struct {
	Twitter  string `json:"twitter"`
	Discord  string `json:"discord"`
	Telegram string `json:"telegram"`
	Github   string `json:"github"`
}

type Vendors struct {
	Image       string   `json:"image"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Socials     Socials  `json:"socials"`
	Offers      []string `json:"offers"`
	Type        string   `json:"type"`
	Hash        string   `json:"hash"`
	Secret      string   `json:"secret"`
}

func MongoCheck(TxHash string, mongoUri string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	var filter bson.D

	collection := client.Database("marketplace").Collection("vendors")

	filter = bson.D{{Key: "tx_hash", Value: TxHash}}

	var vendor Vendors
	err = collection.FindOne(context.TODO(), filter).Decode(&vendor)
	if err == mongo.ErrNoDocuments {
		// Do something when no record found
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	return true
}
