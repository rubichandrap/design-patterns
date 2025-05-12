
package observer

import "fmt"

type Observer interface {
    Update(news string)
}

type NewsAgency struct {
    subscribers []Observer
}

func (n *NewsAgency) AddSubscriber(o Observer) {
    n.subscribers = append(n.subscribers, o)
}

func (n *NewsAgency) Notify(news string) {
    for _, s := range n.subscribers {
        s.Update(news)
    }
}

type Subscriber struct {
    Name string
}

func (s Subscriber) Update(news string) {
    fmt.Printf("%s received news: %s\n", s.Name, news)
}
