package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Message struct {
			Data []byte `json:"data,omitempty"`
			ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

// NewReleasePubSub consumes a Pub/Sub message.
func NewReleasePubSub(w http.ResponseWriter, r *http.Request) {
	// Parse the Pub/Sub message.
	var m PubSubMessage

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