package api

import (
	"config"
	"context"
	"errors"
	"time"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// TODO: Is it possible to have "keyof"/"typeof" with const
// as type?

/*
PageDocument type for MongoDB
*/
type PageDocument struct {
	Title string
	Disabled bool
	Slug string
	Type string
}

// Areas of responsibility (needed for each service)
// 1. Connection to DB
// 2. Collation of data
// 3. Mapping of specific data collations to typed enum services
// 4. Parsing and cleaning of data

/*
StartService initiates a Builder which launches specific service
depending on argument given
*/
func StartService(opts config.Configuration, serviceType string) ([]*PageDocument, error) {

	// Case 1: Should throw error and quit application
	client, err := createMongoConn(opts)

	if err != nil {
		log.Fatalln("Unable to connect to Mongo Instance: ", err)
	}

	getServices, err := availableServices(client, opts)

	if err != nil {
		return nil, err
	}

	// Case 2: If isValidService is incorrect then return error
	if (!isValidService(getServices, serviceType)) {
		return nil, errors.New("Service is not valid")
	}

	data, err := getData(opts, client, serviceType)

	// Case 3: Application should continue, but server will throw 404 since no data found
	if err != nil {
		return nil, err
	}

	return data, nil

}

// Gets the available collections
func availableServices(conn *mongo.Client, config config.Configuration) ([]string, error) {
	opts := options.ListCollections()
	opts.SetNameOnly(true)

	collectionNames, err := conn.Database(config.DB.Name).ListCollectionNames(
		context.Background(), bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	return collectionNames, nil
}

func getData(config config.Configuration, conn *mongo.Client, service string) ([]*PageDocument, error) {

	// Duplicated logic (move to createMongoConn)
	collection := conn.Database(config.DB.Name).Collection(service)
	cur, err := collection.Find(context.Background(), bson.D{})

	defer cur.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	var results []*PageDocument

	for cur.Next(context.TODO()) {
		var elem PageDocument
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func createMongoConn(opts config.Configuration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Cases:
	// 1. MongoDB instance is down (does it panic?)
	// - Connection fails
	// - DB cannot be pinged
	// 2. Load from (.env) -> what if file is not available?
	settings := options.Client().ApplyURI(opts.DB.URL.String())
	client, err := mongo.Connect(ctx, settings)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	return client, nil
}

// TODO: Make recursive
// * Easy to do but not sure about the garbage collection/memory allocation (since it is fixed)
func isValidService(services []string, find string) (bool) {
	for _, value := range  services {
		if value == find {
			return true
		}
	}
	return false
}