package config

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodbURI = "mongodb://%s:%s@%s/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.8.0"

var mongoClient *mongo.Client

func InitMongoClient() error {
	ctx := context.TODO()
	uri := fmt.Sprintf(mongodbURI, viper.GetString("mongo.user"), viper.GetString("mongo.pwd"), viper.GetString("mongo.host"))
	var err error
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	err = mongoClient.Ping(ctx, nil)

	return err
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}
