package server

import (
	"api"
	"config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
func Start(conf config.Configuration) {
	p := conf.Server.Port
	port := fmt.Sprintf(":%s", strconv.Itoa(p))
	http.HandleFunc("/", handleRequests(conf))
	log.Print("Server started on port :", p)
	log.Fatal(http.ListenAndServe(port, nil))
}

/*
handleRequests

Will serve 200 status with JSON. All other GET requests are served with
* 404: For missing data
* 405: For Non-GET requests
* 500: For connection/unrecoverable errors

Fatal errors will quit the server (so check system logs instead)
*/
func handleRequests(options config.Configuration) (func(http.ResponseWriter, *http.Request)) {
	return func (r http.ResponseWriter, s *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Fatal("Program threw exception: ", r)
			}
		}()

		if s.Method != "GET" {
			http.Error(r, http.StatusText(405), 405)
			return
		}

		possibleService := strings.Split(s.RequestURI, "/")[1]
		service, err := api.StartService(options, possibleService);

		if (err == nil) {
			jsonOutput(r, service)
			return
		}
		errorOutput(r, err)
	}
}


func jsonOutput(r http.ResponseWriter, msgObj interface{}) {
	// Case 1: Should create JSON response else throw error
	message, jsonErr := json.Marshal(msgObj)
	if jsonErr != nil {
		errorOutput(r, jsonErr)
		return
	}
	r.Header().Add("Content-Type", "application/json")
	r.Write(message)
}

func errorOutput(r http.ResponseWriter, err error) {
	// Case 1: When error thrown - should output status with correct errorMsg
	log.Println("Error:", err.Error())
	http.Error(r, "Error Occurred", 500)
}