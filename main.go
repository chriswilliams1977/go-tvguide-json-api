package main

/*
"log"
	"net/http"
	"tvguide/routers"
	"context"
	"fmt"
	"log"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
*/
import (
	"log"
	"context"
	"fmt"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

func main() {
	/*
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))*/
	ctx := context.Background()

	// [START fs_initialize]
	// Sets your Google Cloud Platform project ID.
	projectID := "williamscj-serverless-example"

	// Get a Firestore client.
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()
	// [END fs_initialize]

	// [START fs_get_all_users]
	iter := client.Collection("Channels").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
		fmt.Println(doc.Data()["Name"])
		fmt.Println(doc.Data()["ID"])
	}
	// [END fs_get_all_users]
}
