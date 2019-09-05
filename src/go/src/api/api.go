package api

import (
	"fmt"
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
func StartService(service string) {
	switch service {
	case "pages":
		pagesService()
	default:
		fmt.Println("call not available")
	}
}

func pagesService() {
	// createMongoConn()
	fmt.Println("pagesService")
}

// func createMongoConn() {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// }
