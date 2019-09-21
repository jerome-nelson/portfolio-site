package api

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Areas of responsibility (needed for each service)
// 1. Connection to DB
// 2. Collation of data
// 3. Mapping of specific data collations to typed enum services
// 4. Parsing and cleaning of data

/*
StartService initiates a Builder which launches specific service
depending on argument given
*/
func StartService(serviceType string) ([]string, error) {
	// Connection created or err
	client, err := createMongoConn()
	result, dataErr := getData(serviceType)(client)

	if err != nil || dataErr != nil {
		return make([]string, 1), err
	}

	return result, nil

}

func createMongoConn() (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	settings := options.Client().ApplyURI("mongodb://website_admin:website_pass@localhost:27017/website")
	client, err := mongo.Connect(ctx, settings)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	return client, nil
}

func getData(service string) func(*mongo.Client) ([]string, error) {
	return func(client *mongo.Client) ([]string, error) {
		collection := client.Database("website").Collection(service)
		results, err := collection.Find(context.Background(), bson.D{})
		ctx := context.Background()

		if err != nil {
			return nil, err
		}

		defer results.Close(ctx)
		res := formatData(results)
		return res, nil
	}
}

func formatData(results *mongo.Cursor) []string {
	list := make([]string, 0, 10)

	for results.Next(context.Background()) {
		result := results.Current
		// if err != nil {
		// 	list = append(list, err.Error())
		// 	// return err
		// }
		list = append(list, result.String())
	}

	return list
}
