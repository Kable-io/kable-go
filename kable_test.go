package kable_test

import (
	"testing"

	kable "github.com/Kable-io/kable-go"
)

func TestConvert(t *testing.T) {
	c := kable.New()
	c.Event.Record()
	t.Log("Hello World")
	t.Fail()
}
