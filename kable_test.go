package kable_test

import (
	"testing"
	"time"

	"github.com/Kable-io/kable-go"
)

func TestRecord(t *testing.T) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	options := kable.KableOptions{
		KableClientId:     "",
		KableClientSecret: "",
		Debug:             true,
		MaxQueueSize:      10,
	}

	client := kable.New(&options)

	event := kable.Event{
		ClientId:  "elephant-tech",
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"messageId":  "abc123",
			"characters": 12,
			"successful": true,
		},
	}

	client.Record(event)
	// time.Sleep(10 * time.Second)
	// for i := 0; i < 1; i++ {
	// 	c.Record(event)
	// }

	// if err != nil {
	// 	t.Log("Error : ", err)
	// 	t.Fail()
	// }
}
