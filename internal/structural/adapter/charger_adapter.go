
package adapter

import "fmt"

type EuropeanCharger interface {
    ChargeWithTwoPins()
}

type USCharger struct{}
func (u USCharger) ChargeWithThreePins() {
    fmt.Println("Charging with US three-pin plug")
}

type Adapter struct {
    usCharger USCharger
}

func (a Adapter) ChargeWithTwoPins() {
    a.usCharger.ChargeWithThreePins()
    fmt.Println("Adapter converts to two-pin")
}
