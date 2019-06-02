package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tvguide/managers"
	"tvguide/models"
	"strconv"
	"time"

	//routing imports
	"github.com/gorilla/mux"

	//Firestore imports
	"google.golang.org/api/iterator"
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

// handlerFunction for root URL
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our TV Guide V2!")

	// [START fs_get_all_new_releases]
	
	// Parse the Pub/Sub message.
	/*
	var m PubSubMessage

	fmt.Println("message data" , string(m.Message.Data))

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
	}

	name := string(m.Message.Data)
	if name == "" {
			name = "World"
	}
	log.Printf("Hello %s!", name)*/
}

// handlerFunction for /channels/ url path
func HandleChannels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	channels := managers.GetChannelListings()

	json.NewEncoder(w).Encode(channels)
}

// handlerFunction for /channel/{id} url path
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

// handlerFunction for /channel/{id} url path
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

// handlerFunction for /newrelease/ url path
func HandleNewReleases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	// [START fs_initialize]
	// Sets your Google Cloud Platform project ID.
	projectID := "williamscj-serverless-example"

	// Get a Firestore client.
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()
	// [END fs_initialize]

	// [START fs_get_all_new_releases]
	iter := client.Collection("channels").Documents(ctx)
	for {
		
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		
		json.NewEncoder(w).Encode(doc.Data())

	}
	
}

