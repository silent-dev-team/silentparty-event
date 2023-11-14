package sse

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/subscriptions"
)

type Broker struct {
	b *subscriptions.Broker
}

func NewBroker(app core.App) *Broker {
	return &Broker{
		b: app.SubscriptionsBroker(),
	}
}

func (b *Broker) Send(name string, data []byte) error {
	m := subscriptions.Message{
		Name: name,
		Data: data,
	}
	clients := b.b.Clients()

	for _, c := range clients {
		subs := c.Subscriptions()
		for n := range subs {
			if n == name {
				c.Send(m)
				break
			}
		}
	}

	return nil
}
