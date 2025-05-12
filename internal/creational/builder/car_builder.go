
package builder

type Car struct {
    Make  string
    Model string
    Color string
}

type CarBuilder struct {
    car Car
}

func NewCarBuilder() *CarBuilder {
    return &CarBuilder{}
}

func (b *CarBuilder) SetMake(make string) *CarBuilder {
    b.car.Make = make
    return b
}

func (b *CarBuilder) SetModel(model string) *CarBuilder {
    b.car.Model = model
    return b
}

func (b *CarBuilder) SetColor(color string) *CarBuilder {
    b.car.Color = color
    return b
}

func (b *CarBuilder) Build() Car {
    return b.car
}
