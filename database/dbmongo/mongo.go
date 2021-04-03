package dbmongo

import (
	"context"
	"strconv"

	"github.com/open-cmi/goutils/database"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoInit db use mongo
func MongoInit(conf *database.Config) (client *mongo.Client, err error) {
	host := conf.Host
	port := conf.Port
	user := conf.User
	password := conf.Password
	dbname := conf.Database

	uri := "mongodb://" + host + ":" + strconv.Itoa(port)
	if user != "" {
		uri = "mongodb://" + user + ":" + password + "@" + host + ":" + strconv.Itoa(port) + "/" + dbname
	}

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
