package kable

import (
	"fmt"

	"github.com/Kable-io/kable-go/openapi"
)

type Event struct {
	api *openapi.APIClient
}

// TODO : Add type
func (e *Event) Record() {
	fmt.Printf("Recording...")
}
