package hooks

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

func DontAllowIllegalLink(app core.App) func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		app.Dao().ExpandRecord(e.Record, []string{"hp", "ticket"}, nil)
		hp := e.Record.ExpandedOne("hp")
		ticket := e.Record.ExpandedOne("ticket")
		if hp.GetBool("lent") {
			return fmt.Errorf("headphone is already lent")
		}
		if !ticket.GetBool("sold") {
			return fmt.Errorf("ticket is not sold")
		}
		if ticket.GetBool("used") {
			return fmt.Errorf("ticket is already used")
		}
		//TODO: if any link is open, dont allow a new one
		return nil
	}
}

func SetHpToLent(app core.App) func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		hpId := e.Record.GetString("hp")
		ticketId := e.Record.GetString("ticket")
		if hpId == "" || ticketId == "" {
			return nil
		}
		e.Record.Set("start", types.NowDateTime())
		app.Dao().SaveRecord(e.Record)

		hp, err := app.Dao().FindRecordById("hp", hpId)
		if err != nil {
			return err
		}
		hp.Set("lent", true)
		err = app.Dao().SaveRecord(hp)
		if err != nil {
			return err
		}

		ticket, err := app.Dao().FindRecordById("tickets", ticketId)
		if err != nil {
			return err
		}
		ticket.Set("used", true)
		err = app.Dao().SaveRecord(ticket)
		if err != nil {
			return err
		}

		return nil
	}
}
