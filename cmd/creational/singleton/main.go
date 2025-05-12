package main

import "github.com/rubichandrap/design-patterns/internal/creational/factory"

func main() {
	logger := singleton.GetInstance()
	logger.Log("Singleton pattern logger active")
}