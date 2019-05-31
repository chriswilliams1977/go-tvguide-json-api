package managers

import (
	"tvguide/models"
	"time"
	"log"
)

//Return list of channel listings.
func GetChannelListings() []models.Channel {

	channels := []models.Channel{
		models.Channel{
			ID: 1, 
			Listings: []models.Listing {
			models.Listing {Title: "News", Date: "2019-06-07", Time: "07:00:00"},
			models.Listing {Title: "Kids news", Date: "2019-06-07", Time: "07:30:00"},
			models.Listing {Title: "Kids time", Date: "2019-06-07", Time: "08:30:00"},
			models.Listing {Title: "News", Date: "2019-06-07", Time: "09:00:00"},
			models.Listing {Title: "Kids time", Date: "2019-06-07", Time: "10:00:00"},
			models.Listing {Title: "First Dates", Date: "2019-06-07", Time: "10:30:00"},
			models.Listing {Title: "Dating in the Dark", Date: "2019-06-07", Time: "11:00:00"},
			models.Listing {Title: "Talkshow", Date: "2019-06-07", Time: "12:00:00"},
			models.Listing {Title: "Talkshow", Date: "2019-06-07", Time: "12:30:00"},
			models.Listing {Title: "News", Date: "2019-06-07", Time: "13:00:00"},
			models.Listing {Title: "Cartoons", Date: "2019-06-07", Time: "13:30:00"},
			models.Listing {Title: "Cartoons", Date: "2019-06-07", Time: "13:45:00"},
			models.Listing {Title: "News", Date: "2019-06-07", Time: "14:00:00"},
			models.Listing {Title: "News", Date: "2019-06-07", Time: "15:00:00"},
			models.Listing {Title: "News", Date: "2019-06-07", Time: "16:00:00"},
			models.Listing {Title: "Kids time", Date: "2019-06-07", Time: "16:30:00"},
			models.Listing {Title: "Kids news", Date: "2019-06-07", Time: "17:30:00"},
			models.Listing {Title: "News", Date: "2019-06-07", Time: "18:00:00"},
			models.Listing {Title: "Sesame street", Date: "2019-06-07", Time: "19:00:00"},
			models.Listing {Title: "News", Date: "2019-06-07", Time: "20:00:00"},
			models.Listing {Title: "Songfestival", Date: "2019-06-07", Time: "20:30:00"},
			models.Listing {Title: "NewSongfestivals", Date: "2019-06-07", Time: "22:00:00"},
			},
			Name: "Net 1",
		},
		models.Channel{
			ID: 2, 
			Listings: []models.Listing {
			models.Listing {Title: "test3", Date: "2019-07-26", Time: "15:00:00"},
			models.Listing {Title: "test4", Date: "2018-07-13", Time: "11:00:00"},
			},
			Name: "Net 2",
		},

	}
	
	return channels
}


// return channel listings by channel id and time
func GetListingsByChannelId(id int, timeStamp string) models.Channel {

	timeFormat := "15:04:05"
	for _, item := range GetChannelListings() {
		if item.ID == id {
			for i := 0; i < len(item.Listings); i++ {
				
				//parse mux strings to time
				//handles errors
				//get difference between requested time a listing times
				listingTime, err := time.Parse(timeFormat, item.Listings[i].Time)
				if err != nil {
					log.Fatal("listing time format error")
				}
				
				channelRequestedTime, err :=  time.Parse(timeFormat,timeStamp)
				if err != nil {
					log.Fatal("channel time format error")
				}

				if(listingTime.Hour() < channelRequestedTime.Hour()){
					item.Listings = append(item.Listings[:i], item.Listings[i+1:]...)
					i-- // form the remove item index to start iterate next item
				} else {
					if i + 1 < len(item.Listings){
						//if there is another program get index
						nextListingIndex := i + 1
						//get next program time
						nextListingTime, err := time.Parse(timeFormat, item.Listings[nextListingIndex].Time)
						if err != nil{
							log.Fatal("listing time format error")
						}

						//get duration of current program
						programDuration := nextListingTime.Sub(listingTime)
						//if time requested is later than current program duration remove it from listings
						if channelRequestedTime.Sub(listingTime) > programDuration {
							item.Listings = append(item.Listings[:i], item.Listings[i+1:]...)
							i-- // form the remove item index to start iterate next item
						}
					}
				}

			}

			return item
		}
	}

	return models.Channel{}
}
