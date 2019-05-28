package models

//Channel : channel object
type Channel struct {
	ID int
	Listings []Listing
}

//Listing : channel listing object
type Listing struct {
	Title string
	StartDateTime string
}

//Result : return message
type Result struct {
	Message string
}