package main

import (
	"log"
	"net/http"
	"tvguide/routers"
)

func main() {
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
