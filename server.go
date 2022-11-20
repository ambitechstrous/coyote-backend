package main

import (
	"fmt"
	"net/http"
	"strconv"

	"coyote.com/coyote-backend/helpers"
)

func getTimeline(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, "Method is not supported.", http.StatusNotFound)
	}

	userId, err := strconv.ParseUint(request.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		fmt.Fprintf(writer, "Proper user_id must be specified")
	} else {
		jsonTimeline := helpers.GetAggregatedTimeline(userId)
		fmt.Fprintf(writer, jsonTimeline)
	}
}

func main() {
	fmt.Println("Initializing Server at port 8080")
	http.HandleFunc("/timeline", getTimeline)
}
