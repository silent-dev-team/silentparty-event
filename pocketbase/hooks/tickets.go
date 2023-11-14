package hooks

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/silent-dev-team/silentparty-event/pocketbase/utils"
)

func GeneratePin() func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		oldPin := e.Record.Get("_pin")
		fmt.Println("old pin: ", oldPin)
		if oldPin != "" {
			return nil
		}
		pin := utils.GeneratePin()
		e.Record.Set("_pin", pin)
		return nil
	}
}
