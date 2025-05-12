package main

import (
	"github.com/rubichandrap/design-patterns/internal/creational/builder"
	"fmt"
)

func main() {
	car := builder.NewCarBuilder().
		SetMake("Tesla").
		SetModel("Model S").
		SetColor("Black").
		Build()

	fmt.Printf("Built Car: %+v\n", car)
}