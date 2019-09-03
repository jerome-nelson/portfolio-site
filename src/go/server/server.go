package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

const psplit = "/api/*/*"

// const psplit = "api?"

type apiTable map[string]*apiItem
type apiItem struct {
	Path string
}

func main() {
	fmt.Println("[Go API]: Server Started")
	http.HandleFunc("/api/", handleRequests)

	// TODO: Serve from single config for all services
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequests(r http.ResponseWriter, s *http.Request) {
	fmt.Println("[Go API]: Request made")
	fmt.Println(s.RequestURI)
	t := regexp.MustCompile(psplit).Split(s.RequestURI, 3)
	fmt.Println(t)
}
