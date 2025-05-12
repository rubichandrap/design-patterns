package main

import "github.com/rubichandrap/design-patterns/internal/creational/factory"

func main() {
	shape := factory.GetShape("circle")
	if shape != nil {
		shape.Draw()
	}
}