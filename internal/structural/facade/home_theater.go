
package facade

import "fmt"

type DVDPlayer struct{}
func (d DVDPlayer) On() { fmt.Println("DVD Player On") }

type Projector struct{}
func (p Projector) On() { fmt.Println("Projector On") }

type HomeTheaterFacade struct {
    dvd DVDPlayer
    projector Projector
}

func NewHomeTheater() *HomeTheaterFacade {
    return &HomeTheaterFacade{
        dvd: DVDPlayer{},
        projector: Projector{},
    }
}

func (ht *HomeTheaterFacade) WatchMovie() {
    ht.dvd.On()
    ht.projector.On()
    fmt.Println("Enjoy your movie!")
}
