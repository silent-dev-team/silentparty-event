package hooks

import "github.com/pocketbase/pocketbase/core"

func PriceOfPosition(app core.App) func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		app.Dao().ExpandRecord(e.Record, []string{"item"}, nil)
		item := e.Record.ExpandedOne("item")
		e.Record.Set("price", item.GetFloat("price"))
		return nil
	}
}
