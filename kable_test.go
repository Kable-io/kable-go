package kable_test

import (
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"testing"
	"time"

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
		Debug:             true,
		MaxQueueSize:      10,
		BaseUrl:           baseUrl,
	}

	c := kable.New(&options)

	foo := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}

	event := kable.Event{
		ClientId:  "elephant-tech",
		Timestamp: time.Now(),
		Data:      foo,
	}

	res, err := c.EventsApi.Record(event)
	if err != nil {
		t.Log("Error : ", err)
		t.Fail()
	}

	t.Log("Status code : ", res)
}
