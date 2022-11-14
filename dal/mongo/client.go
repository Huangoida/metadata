package mongo

import (
	"context"
	"fmt"
	"metadata/conf"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB *mongo.Database

func InitMangoDb() {
	dsn := fmt.Sprintf(conf.GetConf2().MongoTemplate, conf.GetConf2().Mongo.Username, conf.GetConf2().Mongo.Passwd,
		conf.GetConf2().Mongo.Host, conf.GetConf2().Mongo.Port)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		fmt.Println(err)
	}
	mongoDB = client.Database(conf.GetConf2().Mongo.Database)
}

func GetMongoDb() *mongo.Database {
	return mongoDB
}
