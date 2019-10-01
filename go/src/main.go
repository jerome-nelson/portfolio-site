package main

import (
	"config"
	"flag"
	"log"
	"server"
	"net/url"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("Program threw exception: ", r)
		}
	}()

	options := config.Configuration{}
	var uri string
	flag.IntVar(&options.Server.Port, "serve", -1,"Assign server to port")
	flag.StringVar(&options.DB.Name, "dbname", "", "Name of database")
	flag.StringVar(&uri, "dburl", "", "URL pointing to database instance")
	flag.Parse()

	parsed, err := url.Parse(uri)

	if err != nil || uri == "" {
		log.Fatalln("Database URI parsing error. Did you add -dburl correctly?")
	}

	if (options.DB.Name == "") {
		log.Fatalln("Database parsing error. Did you add --dbname correctly?")
	}

	if (options.Server.Port < 0) {
		log.Fatalln("Please set server port correctly")
	}

	options.DB.URL = parsed
	server.Start(options)
}
