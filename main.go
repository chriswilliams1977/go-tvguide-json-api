package main

/*
"log"
"net/http"
"tvguide/routers"
"context"
"fmt"

"google.golang.org/api/iterator"

"cloud.google.com/go/firestore"
*/
import (
	"log"
	"net/http"
	"tvguide/routers"

	
)

func main() {
	
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
