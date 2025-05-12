
package factory

import "fmt"

type Shape interface {
    Draw()
}

type Circle struct{}
func (c Circle) Draw() {
    fmt.Println("Drawing a Circle")
}

type Square struct{}
func (s Square) Draw() {
    fmt.Println("Drawing a Square")
}

func GetShape(shapeType string) Shape {
    switch shapeType {
    case "circle":
        return Circle{}
    case "square":
        return Square{}
    default:
        return nil
    }
}
