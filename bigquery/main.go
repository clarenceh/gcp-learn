package main

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
)

const (
	projectID = "myek-dev"
)

type Customer struct {
	ID        string
	FirstName string
	LastName  string
}

func main() {
	log.Println("Hello")

	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalln("Fail creating client")
	}

	myDataset := client.Dataset("clarence_test1")

	table := myDataset.Table("customer")

	clarence := &Customer{
		ID:        "1003",
		FirstName: "Simon",
		LastName:  "Yim",
	}

	u := table.Inserter()

	customers := []*Customer{
		clarence,
	}

	if err := u.Put(ctx, customers); err != nil {
		log.Fatalln("Fail inserting record")
	}
}
