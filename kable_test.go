package kable_test

import (
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"github.com/Kable-io/kable-go"
)

func TestRecord(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseUrl, err := url.Parse(os.Getenv("KABLE_BASE_API_URL"))
	if err != nil {
		log.Fatal("Failed to parse KABLE_BASE_API_URL into url: ", err)
	}
	options := kable.KableOptions{
		KableClientId:     os.Getenv("KABLE_CLIENT_ID"),
		KableClientSecret: os.Getenv("KABLE_CLIENT_SECRET"),
		Debug:             false,
		MaxQueueSize:      10,
		BaseUrl:           baseUrl,
	}

	c := kable.New(&options)

	event := kable.Event{
		ClientId:  "elephant-tech",
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"messageId":  "abc123",
			"characters": 12,
			"successful": true,
		},
	}

	c.Record(event)
	time.Sleep(30 * time.Second)

	if err != nil {
		t.Log("Error : ", err)
		t.Fail()
	}
}
