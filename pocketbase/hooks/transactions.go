package hooks

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func TotalPriceOfTransaction(app core.App) func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		app.Dao().ExpandRecord(e.Record, []string{"positions"}, nil)
		positions := e.Record.ExpandedAll("positions")
		e.Record.Set("total", calculateTotal(positions))
		return nil
	}
}

func calculateTotal(positions []*models.Record) float64 {
	total := 0.0
	for _, p := range positions {
		total += p.GetFloat("price") * float64(p.GetInt("amount"))
	}
	return total
}
