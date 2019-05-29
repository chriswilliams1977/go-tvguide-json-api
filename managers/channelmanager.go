package managers

import (
	"tvguide/models"
)

//Return list of channel listings.
func GetChannelListings() []models.Channel {

	channels := []models.Channel{
		models.Channel{
			ID: 1, 
			Listings: []models.Listing {
			models.Listing {Title: "test1", StartDateTime: "2019-07-26 15:00:00"},
			models.Listing {Title: "test2", StartDateTime: "2018-07-13 11:00:00"},
			},
		},
		models.Channel{
			ID: 2, 
			Listings: []models.Listing {
			models.Listing {Title: "test3", StartDateTime: "2019-07-26 15:00:00"},
			models.Listing {Title: "test4", StartDateTime: "2018-07-13 11:00:00"},
			},
		},

	}
	return channels
}

// return user by ID
func GetListingsByChannelId(id int) models.Channel {

	for _, item := range GetChannelListings() {
		if item.ID == id {
			return item
		}
	}

	return models.Channel{}
}
