package sse

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/silent-dev-team/silentparty-event/pocketbase/measurements"
	m "github.com/silent-dev-team/silentparty-event/pocketbase/models"
)

func OnSubscription(app core.App) func(e *core.RealtimeSubscribeEvent) error {
	return func(e *core.RealtimeSubscribeEvent) error {
		SendOnSubscription(e, "userstats", func() (any, error) {
			stats, _ := measurements.GetUserStats(app)
			return stats, nil
		})
		SendOnSubscription(e, "djs", func() (any, error) {
			return m.GetAllDjs(app)
		})
		SendOnSubscription(e, "marquee", func() (any, error) {
			return m.GetAllMarques(app)
		})
		SendOnSubscription(e, "shop_items", func() (any, error) {
			return m.GetAllShopItems(app)
		})
		return nil
	}
}

func OnMessage(app core.App) func(e *core.RealtimeMessageEvent) error {
	return func(e *core.RealtimeMessageEvent) error {
		SendOnMessage(e, "userstats", func() (any, error) {
			stats, _ := measurements.GetUserStats(app)
			return stats, nil
		})
		SendOnMessage(e, "marquee", func() (any, error) {
			return m.GetAllMarques(app)
		})
		SendOnMessage(e, "djs", func() (any, error) {
			return m.GetAllDjs(app)
		})
		SendOnMessage(e, "shop_items", func() (any, error) {
			return m.GetAllShopItems(app)
		})
		return nil
	}
}
