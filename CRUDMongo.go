package main

import (
	"basic_crud/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insertOneIntoPodcast(tempPodcast model.Podcast, podcastsCollection *mongo.Collection) {
	//podcastsCollection := quickstartDatabase.Collection("podcasts")

	fmt.Println(tempPodcast)
	result, err := podcastsCollection.InsertOne(context.TODO(), tempPodcast)
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted a single document", result.InsertedID)
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")

	tempPodcast := model.Podcast{
		Title:  "English",
		Author: "Tamim Tamim",
		Tags:   []string{"Efg", "Hike"},
	}
	insertOneIntoPodcast(tempPodcast, podcastsCollection)

}
