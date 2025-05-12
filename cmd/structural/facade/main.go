package main

import (
    "github.com/rubichandrap/design-patterns/internal/structural/facade"
)

func main() {
    homeTheater := facade.NewHomeTheaterFacade()

	homeTheater.WatchMovie("Inception")
    homeTheater.EndMovie()
}