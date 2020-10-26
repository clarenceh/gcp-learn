package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type Example struct {
	id string
}

func main() {
	log.Println("Hello")

	event := cloudevents.NewEvent()

	//event.SetSpecVersion("1.0")
	event.SetSource("myek-api")
	event.SetType("Test")
	event.SetID("abc123")
	event.SetTime(time.Now())
	event.SetData("application/json", &Example{
		id: "ABD",
	})

	reqBodyBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Fail marshall event: %s", err)
		os.Exit(1)
	}

	reqBody := bytes.NewReader(reqBodyBytes)

	resp, err := http.Post("http://localhost:8080", "application/json", reqBody)
	if err != nil {
		log.Fatalln("Fail posting request")
		os.Exit(1)
	}

	log.Printf("Response: %+v", resp.Body)
}
