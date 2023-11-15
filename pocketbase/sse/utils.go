package sse

import (
	"encoding/json"
	"slices"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/subscriptions"
)

func SendOnSubscription(e *core.RealtimeSubscribeEvent, name string, f func() (any, error)) error {
	if slices.Contains(e.Subscriptions, name) {
		data, err := f()
		if err != nil {
			return err
		}
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}
		m := subscriptions.Message{
			Name: name,
			Data: b,
		}
		e.Client.Send(m)
		return nil
	}
	return nil
}

func SendOnMessage(e *core.RealtimeMessageEvent, name string, f func() (any, error)) error {
	if e.Message.Name == name {
		data, err := f()
		if err != nil {
			return err
		}
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}
		e.Message.Data = b
		return nil
	}
	return nil
}
