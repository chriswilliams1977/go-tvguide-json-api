package managers

import (
	"tvguide/models"
	"time"
	"log"
	"fmt"
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
			models.Listing {Title: "News", Date: "2019-07-13", Time: "07:00:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "07:30:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "08:30:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "09:00:00"},
			models.Listing {Title: "Soap Series", Date: "2018-07-13", Time: "10:00:00"},
			models.Listing {Title: "Talkshow", Date: "2018-07-13", Time: "10:30:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "11:00:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "12:00:00"},
			models.Listing {Title: "Talkshow", Date: "2018-07-13", Time: "12:30:00"},
			models.Listing {Title: "Boardwalk Empire", Date: "2018-07-13", Time: "13:30:00"},
			models.Listing {Title: "BNN at Work", Date: "2018-07-13", Time: "07:30:00"},
			models.Listing {Title: "Cartoons", Date: "2018-07-13", Time: "13:45:00"},
			models.Listing {Title: "Cartoons", Date: "2018-07-13", Time: "14:00:00"},
			models.Listing {Title: "Talkshow", Date: "2018-07-13", Time: "15:00:00"},
			models.Listing {Title: "Veronica Mars", Date: "2018-07-13", Time: "16:00:00"},
			models.Listing {Title: "Cash Cab", Date: "2018-07-13", Time: "16:30:00"},
			models.Listing {Title: "Kids Time", Date: "2018-07-13", Time: "17:30:00"},
			models.Listing {Title: "Kids News", Date: "2018-07-13", Time: "18:00:00"},
			models.Listing {Title: "Kids News", Date: "2018-07-13", Time: "19:00:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "20:00:00"},
			models.Listing {Title: "Soap Series", Date: "2018-07-13", Time: "20:30:00"},
			models.Listing {Title: "Flashback", Date: "2018-07-13", Time: "22:00:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "00:00:00"},
			},
			Name: "Net 2",
		},
		models.Channel{
			ID: 3, 
			Listings: []models.Listing {
			models.Listing {Title: "Cartoons", Date: "2019-07-13", Time: "07:00:00"},
			models.Listing {Title: "Cartoons", Date: "2018-07-13", Time: "07:30:00"},
			models.Listing {Title: "Cartoons", Date: "2018-07-13", Time: "08:30:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "09:00:00"},
			models.Listing {Title: "Kids Time", Date: "2018-07-13", Time: "10:00:00"},
			models.Listing {Title: "Cartoons", Date: "2018-07-13", Time: "10:30:00"},
			models.Listing {Title: "First Dates Hotel", Date: "2018-07-13", Time: "11:00:00"},
			models.Listing {Title: "First Dates", Date: "2018-07-13", Time: "12:00:00"},
			models.Listing {Title: "Talkshow", Date: "2018-07-13", Time: "13:00:00"},
			models.Listing {Title: "Max TV", Date: "2018-07-13", Time: "13:30:00"},
			models.Listing {Title: "Max TV", Date: "2018-07-13", Time: "13:45:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "14:00:00"},
			models.Listing {Title: "Pretty & Single", Date: "2018-07-13", Time: "15:00:00"},
			models.Listing {Title: "Taxi", Date: "2018-07-13", Time: "16:00:00"},
			models.Listing {Title: "Comedy Hour", Date: "2018-07-13", Time: "16:30:00"},
			models.Listing {Title: "Comedy Hour", Date: "2018-07-13", Time: "17:30:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "18:00:00"},
			models.Listing {Title: "Sports Hour", Date: "2018-07-13", Time: "19:00:00"},
			models.Listing {Title: "Movie", Date: "2018-07-13", Time: "20:00:00"},
			models.Listing {Title: "Movie", Date: "2018-07-13", Time: "22:00:00"},
			},
			Name: "Net 3",
		},
		models.Channel{
			ID: 4, 
			Listings: []models.Listing {
			models.Listing {Title: "News", Date: "2019-07-13", Time: "07:00:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "07:30:00"},
			models.Listing {Title: "Telekids", Date: "2018-07-13", Time: "08:30:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "09:00:00"},
			models.Listing {Title: "Bold and the Beautiful", Date: "2018-07-13", Time: "10:00:00"},
			models.Listing {Title: "House Vision", Date: "2018-07-13", Time: "10:30:00"},
			models.Listing {Title: "House Vision", Date: "2018-07-13", Time: "11:00:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "12:00:00"},
			models.Listing {Title: "The Good Doctor", Date: "2018-07-13", Time: "12:30:00"},
			models.Listing {Title: "Telekids", Date: "2018-07-13", Time: "13:00:00"},
			models.Listing {Title: "Telekids", Date: "2018-07-13", Time: "13:30:00"},
			models.Listing {Title: "Telekids", Date: "2018-07-13", Time: "13:45:00"},
			models.Listing {Title: "Telekids", Date: "2018-07-13", Time: "14:00:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "15:00:00"},
			models.Listing {Title: "Talkshow", Date: "2018-07-13", Time: "16:00:00"},
			models.Listing {Title: "RTL Boulevard", Date: "2018-07-13", Time: "16:30:00"},
			models.Listing {Title: "The Blacklist", Date: "2018-07-13", Time: "18:00:00"},
			models.Listing {Title: "All you need is love", Date: "2018-07-13", Time: "19:00:00"},
			models.Listing {Title: "News", Date: "2018-07-13", Time: "20:00:00"},
			models.Listing {Title: "The Bachelor", Date: "2018-07-13", Time: "20:30:00"},
			models.Listing {Title: "Marriage at First Sight", Date: "2018-07-13", Time: "22:00:00"},
			models.Listing {Title: "Movie", Date: "2018-07-13", Time: "00:00:00"},
			},
			Name: "RTL 4",
		},
		models.Channel{
			ID: 5, 
			Listings: []models.Listing {
			models.Listing {Title: "Star Wars A New Hope", Date: "2019-07-13", Time: "16:00:00"},
			models.Listing {Title: "Batman begins", Date: "2018-07-13", Time: "18:00:00"},
			models.Listing {Title: "Avengers Endgame", Date: "2018-07-13", Time: "20:30:00"},
			models.Listing {Title: "Deadpool", Date: "2018-07-13", Time: "22:00:00"},
			models.Listing {Title: "Bridemaids", Date: "2018-07-13", Time: "00:00:00"},
			},
			Name: "Movie Channel",
		},
		models.Channel{
			ID: 6, 
			Listings: []models.Listing {
			models.Listing {Title: "Basketball", Date: "2019-07-13", Time: "09:00:00"},
			models.Listing {Title: "Basketball", Date: "2018-07-13", Time: "10:00:00"},
			models.Listing {Title: "Extreme Sports", Date: "2018-07-13", Time: "10:30:00"},
			models.Listing {Title: "Volleyball", Date: "2018-07-13", Time: "11:00:00"},
			models.Listing {Title: "Hockey", Date: "2018-07-13", Time: "12:00:00"},
			models.Listing {Title: "Soccer", Date: "2018-07-13", Time: "12:30:00"},
			models.Listing {Title: "Tennis", Date: "2018-07-13", Time: "13:00:00"},
			models.Listing {Title: "Tennis", Date: "2018-07-13", Time: "13:30:00"},
			models.Listing {Title: "Formula 1", Date: "2018-07-13", Time: "13:45:00"},
			models.Listing {Title: "Formula 1", Date: "2018-07-13", Time: "14:00:00"},
			models.Listing {Title: "Extreme Sports", Date: "2018-07-13", Time: "15:00:00"},
			models.Listing {Title: "NBA", Date: "2018-07-13", Time: "16:00:00"},
			models.Listing {Title: "Soccer News", Date: "2018-07-13", Time: "16:30:00"},
			models.Listing {Title: "NBA", Date: "2018-07-13", Time: "17:30:00"},
			models.Listing {Title: "Baseball", Date: "2018-07-13", Time: "18:00:00"},
			models.Listing {Title: "Baseball", Date: "2018-07-13", Time: "19:00:00"},
			models.Listing {Title: "Soccer", Date: "2018-07-13", Time: "20:00:00"},
			models.Listing {Title: "Soccer", Date: "2018-07-13", Time: "20:30:00"},
			models.Listing {Title: "Soccer", Date: "2018-07-13", Time: "22:00:00"},
			models.Listing {Title: "Soccer", Date: "2018-07-13", Time: "00:00:00"},
			},
			Name: "Sports Channel",
		},
		models.Channel{
			ID: 7, 
			Listings: []models.Listing {
			models.Listing {Title: "The Simpsons", Date: "2019-07-13", Time: "10:00:00"},
			models.Listing {Title: "Family Guy", Date: "2018-07-13", Time: "10:30:00"},
			models.Listing {Title: "Rick & Morty", Date: "2018-07-13", Time: "11:00:00"},
			models.Listing {Title: "Southpark", Date: "2018-07-13", Time: "12:00:00"},
			models.Listing {Title: "Modern Family", Date: "2018-07-13", Time: "12:30:00"},
			models.Listing {Title: "How I met your mother", Date: "2018-07-13", Time: "13:00:00"},
			models.Listing {Title: "The Office", Date: "2018-07-13", Time: "13:30:00"},
			models.Listing {Title: "Young Sheldon", Date: "2018-07-13", Time: "13:45:00"},
			models.Listing {Title: "The Big Bang Theory", Date: "2018-07-13", Time: "14:00:00"},
			models.Listing {Title: "Mom", Date: "2018-07-13", Time: "15:00:00"},
			models.Listing {Title: "Rick & Morty", Date: "2018-07-13", Time: "16:00:00"},
			models.Listing {Title: "Celebrity Standup", Date: "2018-07-13", Time: "16:30:00"},
			models.Listing {Title: "Celebrity Standup", Date: "2018-07-13", Time: "17:30:00"},
			models.Listing {Title: "The Simpsons", Date: "2018-07-13", Time: "18:00:00"},
			models.Listing {Title: "Family Guy", Date: "2018-07-13", Time: "19:00:00"},
			models.Listing {Title: "Rick & Morty", Date: "2018-07-13", Time: "20:00:00"},
			models.Listing {Title: "Southpark", Date: "2018-07-13", Time: "20:30:00"},
			models.Listing {Title: "Comedy Roast", Date: "2018-07-13", Time: "22:00:00"},
			models.Listing {Title: "The Daily Show", Date: "2018-07-13", Time: "00:00:00"},
			},
			Name: "Comedy Central",
		},
		models.Channel{
			ID: 8, 
			Listings: []models.Listing {
			models.Listing {Title: "Adventure time", Date: "2019-07-13", Time: "07:00:00"},
			models.Listing {Title: "Dexters Lab", Date: "2018-07-13", Time: "07:30:00"},
			models.Listing {Title: "Powerpuff Girls", Date: "2018-07-13", Time: "08:30:00"},
			models.Listing {Title: "Spider-man", Date: "2018-07-13", Time: "09:00:00"},
			models.Listing {Title: "Batman", Date: "2018-07-13", Time: "10:00:00"},
			models.Listing {Title: "Johnny Bravo", Date: "2018-07-13", Time: "10:30:00"},
			models.Listing {Title: "Dexters Lab", Date: "2018-07-13", Time: "11:00:00"},
			models.Listing {Title: "Adventure time", Date: "2018-07-13", Time: "12:00:00"},
			models.Listing {Title: "Samurai Jack", Date: "2018-07-13", Time: "12:30:00"},
			models.Listing {Title: "Ed Edd and Eddy", Date: "2018-07-13", Time: "13:00:00"},
			models.Listing {Title: "Courage the cowardly dog", Date: "2018-07-13", Time: "13:30:00"},
			models.Listing {Title: "Tom & Jerry", Date: "2018-07-13", Time: "13:45:00"},
			models.Listing {Title: "Tom & Jerry", Date: "2018-07-13", Time: "14:00:00"},
			models.Listing {Title: "Powerpuff Girls", Date: "2018-07-13", Time: "15:00:00"},
			models.Listing {Title: "Dexters Lab", Date: "2018-07-13", Time: "16:00:00"},
			models.Listing {Title: "Ed Edd and Eddy", Date: "2018-07-13", Time: "16:30:00"},
			models.Listing {Title: "Samurai Jack", Date: "2018-07-13", Time: "17:30:00"},
			models.Listing {Title: "Adventure time", Date: "2018-07-13", Time: "18:00:00"},
			models.Listing {Title: "Jonny Bravo", Date: "2018-07-13", Time: "19:00:00"},
			models.Listing {Title: "Family Guy", Date: "2018-07-13", Time: "20:00:00"},
			models.Listing {Title: "Batman", Date: "2018-07-13", Time: "20:30:00"},
			models.Listing {Title: "Spider-man", Date: "2018-07-13", Time: "22:00:00"},
			models.Listing {Title: "Jonny Bravo", Date: "2018-07-13", Time: "00:00:00"},
			},
			Name: "Cartoon Network",
		},
		models.Channel{
			ID: 9, 
			Listings: []models.Listing {
			models.Listing {Title: "Life below zero", Date: "2019-07-13", Time: "07:00:00"},
			models.Listing {Title: "Life below zero", Date: "2018-07-13", Time: "07:30:00"},
			models.Listing {Title: "Life below zero", Date: "2018-07-13", Time: "08:30:00"},
			models.Listing {Title: "Genious", Date: "2018-07-13", Time: "09:00:00"},
			models.Listing {Title: "Genious", Date: "2018-07-13", Time: "10:00:00"},
			models.Listing {Title: "Genious", Date: "2018-07-13", Time: "10:30:00"},
			models.Listing {Title: "Megastructures", Date: "2018-07-13", Time: "11:00:00"},
			models.Listing {Title: "Street Genious", Date: "2018-07-13", Time: "12:00:00"},
			models.Listing {Title: "Dog Whisperer", Date: "2018-07-13", Time: "12:30:00"},
			models.Listing {Title: "Dog Town", Date: "2018-07-13", Time: "13:00:00"},
			models.Listing {Title: "Wild", Date: "2018-07-13", Time: "13:30:00"},
			models.Listing {Title: "Wild", Date: "2018-07-13", Time: "13:45:00"},
			models.Listing {Title: "Cosmos", Date: "2018-07-13", Time: "14:00:00"},
			models.Listing {Title: "Hard Time", Date: "2018-07-13", Time: "15:00:00"},
			models.Listing {Title: "Wicked Tuna", Date: "2018-07-13", Time: "16:00:00"},
			models.Listing {Title: "Be the creature", Date: "2018-07-13", Time: "16:30:00"},
			models.Listing {Title: "Megastructures", Date: "2018-07-13", Time: "17:30:00"},
			models.Listing {Title: "Street Genious", Date: "2018-07-13", Time: "18:00:00"},
			models.Listing {Title: "Genious", Date: "2018-07-13", Time: "19:00:00"},
			models.Listing {Title: "Life below zero", Date: "2018-07-13", Time: "20:00:00"},
			models.Listing {Title: "Wild", Date: "2018-07-13", Time: "20:30:00"},
			models.Listing {Title: "Wild", Date: "2018-07-13", Time: "22:00:00"},
			models.Listing {Title: "Wild", Date: "2018-07-13", Time: "00:00:00"},
			},
			Name: "National Geographic",
		},
		models.Channel{
			ID: 10, 
			Listings: []models.Listing {
			models.Listing {Title: "Videoclips", Date: "2019-07-13", Time: "07:00:00"},
			models.Listing {Title: "Videoclips", Date: "2018-07-13", Time: "07:30:00"},
			models.Listing {Title: "Videoclips", Date: "2018-07-13", Time: "08:30:00"},
			models.Listing {Title: "Videoclips", Date: "2018-07-13", Time: "09:00:00"},
			models.Listing {Title: "Videoclips", Date: "2018-07-13", Time: "10:00:00"},
			models.Listing {Title: "The Real World", Date: "2018-07-13", Time: "10:30:00"},
			models.Listing {Title: "Catfish Marathon", Date: "2018-07-13", Time: "11:00:00"},
			models.Listing {Title: "Catfish Marathon", Date: "2018-07-13", Time: "12:00:00"},
			models.Listing {Title: "Pimp my ride", Date: "2018-07-13", Time: "12:30:00"},
			models.Listing {Title: "Jersey Shore", Date: "2018-07-13", Time: "13:00:00"},
			models.Listing {Title: "Jersey Shore", Date: "2018-07-13", Time: "13:30:00"},
			models.Listing {Title: "Daria", Date: "2018-07-13", Time: "13:45:00"},
			models.Listing {Title: "The Real World", Date: "2018-07-13", Time: "14:00:00"},
			models.Listing {Title: "The Osbournes", Date: "2018-07-13", Time: "15:00:00"},
			models.Listing {Title: "Teenwolf", Date: "2018-07-13", Time: "16:00:00"},
			models.Listing {Title: "MTV Unplugged", Date: "2018-07-13", Time: "16:30:00"},
			models.Listing {Title: "Rupauls Drag Race", Date: "2018-07-13", Time: "17:30:00"},
			models.Listing {Title: "Ridiculousness", Date: "2018-07-13", Time: "18:00:00"},
			models.Listing {Title: "Punk'd", Date: "2018-07-13", Time: "19:00:00"},
			models.Listing {Title: "Jersey Shore", Date: "2018-07-13", Time: "20:00:00"},
			models.Listing {Title: "MTV Awards", Date: "2018-07-13", Time: "20:30:00"},
			models.Listing {Title: "Beavis & Butthead", Date: "2018-07-13", Time: "22:00:00"},
			models.Listing {Title: "Videoclips", Date: "2018-07-13", Time: "00:00:00"},
			},
			Name: "MTV",
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
					log.Fatal("listing time format errors")
				}
				
				channelRequestedTime, err :=  time.Parse(timeFormat,timeStamp)
				if err != nil {
					log.Fatal("channel time format error")
				}
				fmt.Println("Listing time is: ", listingTime)
				fmt.Println("Channel requested time is: ", channelRequestedTime)
				if (listingTime.Hour() < channelRequestedTime.Hour()){
					fmt.Println("there are listings left")
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
						fmt.Println("Next listing item ", nextListingTime)
						fmt.Println("Program duration ", programDuration)
						if(listingTime.Hour() == channelRequestedTime.Hour()) {
							//if time requested is later than current program duration remove it from listings
							if channelRequestedTime.Sub(listingTime) > programDuration {
							item.Listings = append(item.Listings[:i], item.Listings[i+1:]...)
							i-- // form the remove item index to start iterate next item
							}
						}
						
					}
				}

			}

			return item
		}
	}

	return models.Channel{}
}
