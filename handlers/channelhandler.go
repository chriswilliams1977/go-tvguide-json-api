package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tvguide/managers"
	"tvguide/models"
	"strconv"
	"time"
	"log"

	"github.com/gorilla/mux"
)

// handlerFunction for root URL
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our TV Guide V2!")
	
	// Parse the Pub/Sub message.
	var m PubSubMessage

	fmt.Println("message data" , m.Message.Data)

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
	}

	name := string(m.Message.Data)
	if name == "" {
			name = "World"
	}
	log.Printf("Hello %s!", name)
}

// handlerFunction for /channels/ url path
func HandleChannels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channels := managers.GetChannelListings()

	json.NewEncoder(w).Encode(channels)
}

// handlerFunction for /user/{id} url path
func HandleChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	channelId, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	t := time.Now()
	channelTime := t.Format("15:04:05")

	switch r.Method {

	case "GET":

		channelListing := managers.GetListingsByChannelId(channelId, channelTime)
		json.NewEncoder(w).Encode(channelListing)

	case "DELETE":
		json.NewEncoder(w).Encode(models.Result{
			Message: fmt.Sprintf("HTTP %s Method", r.Method),
		})

	case "POST":
		json.NewEncoder(w).Encode(models.Result{
			Message: fmt.Sprintf("HTTP %s Method ", r.Method),
		})
	}
}

// handlerFunction for /user/{id} url path
func HandleChannelTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	channelId, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	channelTime := vars["time"]	
	if channelTime != "" {
		w.WriteHeader(http.StatusNotFound)
	}

	switch r.Method {

	case "GET":

		channelListing := managers.GetListingsByChannelId(channelId, channelTime)
		json.NewEncoder(w).Encode(channelListing)

	case "DELETE":
		json.NewEncoder(w).Encode(models.Result{
			Message: fmt.Sprintf("HTTP %s Method", r.Method),
		})

	case "POST":
		json.NewEncoder(w).Encode(models.Result{
			Message: fmt.Sprintf("HTTP %s Method ", r.Method),
		})
	}
}