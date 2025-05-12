package main

import (
	"github.com/rubichandrap/design-patterns/internal/behavioral/observer"
)

func main() {
	agency := observer.NewsAgency{}

	s1 := observer.Subscriber{Name: "Alice"}
	s2 := observer.Subscriber{Name: "Bob"}

	agency.AddSubscriber(s1)
	agency.AddSubscriber(s2)

	agency.Notify("Go 1.21 Released!")
}