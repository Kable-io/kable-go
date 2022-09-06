package kable

import "fmt"

// import (
// 	"github.com/Kable-io/kable-go/internal/openapi"
// )

type (
	KableOptions struct {
	}

	Kable struct {
		Event *Event
	}
)

func (p Kable) Foo() {
	fmt.Println("Hello World!")
}

func New() *Kable {
	return &Kable{}
}
