package kable_test

import (
	"log"
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

	options := kable.KableOptions{
		KableClientId:     os.Getenv("KABLE_CLIENT_ID"),
		KableClientSecret: os.Getenv("KABLE_CLIENT_SECRET"),
		Debug:             true,
		MaxQueueSize:      10,
		BaseUrl:           os.Getenv("KABLE_BASE_API_URL"),
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
