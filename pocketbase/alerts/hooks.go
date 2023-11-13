package alerts

import (
	"fmt"
	"strings"

	"github.com/pocketbase/pocketbase/core"
)

func (h *Handler) GetAfterCreatedAlertHook() func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		m := e.Record.GetString("msg")
		from := strings.ToUpper(e.Record.GetString("from"))
		text := fmt.Sprintf("⚠️ %s ⚠️\n\n%s", from, m)
		message, _ := h.bot.SendGroupMessage("default", text)
		e.Record.Set("tg_msg_id", message.MessageID)
		h.pb.Dao().SaveRecord(e.Record)
		h.AddAlert(NewAlertFromRecord(e.Record))
		return nil
	}
}

func (h *Handler) GetAfterUpdatedAlertHook() func(e *core.RecordUpdateEvent) error {
	return func(e *core.RecordUpdateEvent) error {
		from := strings.ToUpper(e.Record.GetString("from"))
		active := e.Record.GetBool("active")
		if active {
			return nil
		}
		text := fmt.Sprintf("✅ %s ✅\n\n Hat sich erledigt", from)
		h.bot.EditGroupMessage("default", e.Record.GetInt("tg_msg_id"), text)
		h.RemoveAlert(e.Record.Id)
		return nil
	}
}
