package main

import (
	"flag"
	"fmt"
	"server"
	"strconv"
)

var flagPort int

func init() {
	flag.IntVar(&flagPort, "port", 1080, "port assignment for go application")
	flag.Parse()
}

func main() {
	// TODO: Add command line arguments
	// TODO: Serve from single config for all services
	// TODO: Make more perfomant concatenation
	p := strconv.Itoa(flagPort)
	r := fmt.Sprintf(":%s", p)
	server.StartFactory(r)
}
