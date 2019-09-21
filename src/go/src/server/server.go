package server

import (
	"api"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

/*
 Response payload
*/
type responsePayload struct {
	Response []string
	ErrorMsg string
}

/*
StartFactory function which launches singleton server instance
*/
func StartFactory(port string) {
	http.HandleFunc("/api/", handleRequests)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleRequests(r http.ResponseWriter, s *http.Request) {
	t := regexp.MustCompile("/api/*").Split(s.RequestURI, 5)
	payload := responsePayload{}
	u, err := api.StartService(t[1])

	if err != nil {
		r.WriteHeader(404)
		payload.ErrorMsg = err.Error()
	}

	payload.Response = u
	message, jsonErr := json.Marshal(payload)

	if jsonErr != nil {
		r.WriteHeader(500)
		message = []byte("JSON Compilation Error:" + jsonErr.Error())
	}

	r.Header().Add("Content-Type", "application/json")
	r.Write(message)
}
