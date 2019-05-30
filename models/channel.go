package models

//Channel : channel object
type Channel struct {
	ID int
	Listings []Listing
	Name string
}

//Listing : channel listing object
type Listing struct {
	Title string
	Date string
	Time string
}

//Result : return message
type Result struct {
	Message string
}