package main

import (
    "github.com/rubichandrap/design-patterns/internal/structural/adapter"
)

func main() {
    usCharger := adapter.USCharger{}
    europeanAdapter := adapter.Adapter{usCharger: usCharger}

	europeanAdapter.ChargeWithTwoPins()
}