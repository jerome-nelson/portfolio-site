package main

import (
	"flag"
	"fmt"
	"server"
	"strconv"
)

var flagPort int

func init() {
	flag.IntVar(&flagPort, "port", -1, "port assignment for go application")
	flag.Parse()
}

func main() {
	p := strconv.Itoa(flagPort)
	r := fmt.Sprintf(":%s", p)
	server.StartFactory(r)
}
