package server

import (
	"api"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// Cases to cover
// 1. Handling error (if service doesn't exist)
// 2. if port is busy throw exception
// 3.

/*
StartFactory function which launches server
*/
func StartFactory(port string) {
	fmt.Println("[Server]: server launched")
	http.HandleFunc("/api/", handleRequests)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Cases to cover
// 1. api/pages/ - returns service
// 2. /api/pages/23123 - returns service
func handleRequests(r http.ResponseWriter, s *http.Request) {
	t := regexp.MustCompile("/api/*").Split(s.RequestURI, 5)
	api.StartService(t[1])
}
