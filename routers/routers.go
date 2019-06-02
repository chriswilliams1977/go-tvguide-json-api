package routers

import (
	"net/http"
	"tvguide/handlers"

	"github.com/gorilla/mux"
)

//Route : route object
type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes []Route

// init method which will be automatically run 
func init() {

	// init root URL and handler function  
	routes = append(routes, Route{
		Name:        "Index",
		Methods:     []string{"GET"},
		Pattern:     "/",
		HandlerFunc: handlers.Index,
	})

	/* init /channel/{id} URL and handler function. 
	One handler function for three type of requests */
	routes = append(routes, Route{
		Name:        "Channel",
		Methods:     []string{"DELETE", "GET", "POST"},
		Pattern:     "/channel/{id}",
		HandlerFunc: handlers.HandleChannel,
	})

	/* init /channels URL and handler function. */
	routes = append(routes, Route{
		Name:        "Channels",
		Methods:     []string{"GET"},
		Pattern:     "/channels",
		HandlerFunc: handlers.HandleChannels,
	})

	/* init /channel/{id}/{time} URL and handler function. */
	routes = append(routes, Route{
		Name:        "ChannelTime",
		Methods:     []string{"DELETE", "GET", "POST"},
		Pattern:     "/channel/{id}/{time}",
		HandlerFunc: handlers.HandleChannelTime,
	})

	/* init /newrleases URL and handler function. */
	routes = append(routes, Route{
		Name:        "ChannelReleases",
		Methods:     []string{"DELETE", "GET", "POST"},
		Pattern:     "/newreleases/",
		HandlerFunc: handlers.HandleNewReleases,
	})
}

/* registering all url paths */
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Methods...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}