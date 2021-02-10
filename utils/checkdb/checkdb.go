package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	file, err := os.Open("unique.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//initialise mongo client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 999999999*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	collection := client.Database("LazyXSS").Collection("searchedURLs")
	time.Sleep(5 * time.Second)

	//go through each line in a file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var url string = scanner.Text()

		count, err := collection.CountDocuments(ctx, bson.M{"Found": url})
		if err != nil {
			log.Fatal(err)
		}
		if count >= 1 {
		} else {
			fmt.Println(url)
		}
	}
}
