package handlers


// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Message struct {
			Data []byte `json:"data,omitempty"`
			ID   string `json:"messageId"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}