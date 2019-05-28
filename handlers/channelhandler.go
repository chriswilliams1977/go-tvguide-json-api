package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tvguide/managers"
	"tvguide/models"
	"strconv"

	"github.com/gorilla/mux"
)

// handlerFunction for root URL
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our TV Guide!")
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

	switch r.Method {

	case "GET":

		channelListing := managers.GetListingsByChannelId(channelId)
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