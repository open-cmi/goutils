package dbmongo

import (
	"context"
	"strconv"

	"github.com/open-cmi/goutils/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

// MongoInit db use mongo
func MongoInit() error {
	host := config.Conf.GetStringMap("model")["host"].(string)
	port := config.Conf.GetStringMap("model")["port"].(int)
	user := config.Conf.GetStringMap("model")["user"].(string)
	password := config.Conf.GetStringMap("model")["password"].(string)
	dbname := config.Conf.GetStringMap("model")["dbname"].(string)

	uri := "mongodb://" + host + ":" + strconv.Itoa(port)
	if user != "" {
		uri = "mongodb://" + user + ":" + password + "@" + host + ":" + strconv.Itoa(port) + "/" + dbname
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}

	MongoClient = client
	return nil
}

// MongoFini close connection
func MongoFini() {
	MongoClient.Disconnect(context.TODO())
	return
}
