package models

//Channel : channel object
type Channel struct {
	ID int `json:"Channel"`
	Listings []Listing `json:"Listings"`
	Name string `json:"ChannelName"`
}

//Listing : channel listing object
type Listing struct {
	Title string `json:"Title"`
	StartDateTime string `json:"StartDateTime"`
	//Time string `json:"Listings"`
}

//Result : return message
type Result struct {
	Message string
}