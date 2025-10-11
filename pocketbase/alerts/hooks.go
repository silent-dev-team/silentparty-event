package alerts

import (
	"fmt"
	"log"
	"strings"

	"github.com/pocketbase/pocketbase/core"
)

func (h *Handler) GetAfterCreatedAlertHook() func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		log.Println("New alert:", e.Record)
		channel := e.Record.GetString("channel")
		m := e.Record.GetString("msg")
		from := strings.ToUpper(e.Record.GetString("from"))
		text := fmt.Sprintf("⚠️ %s ⚠️\n\n%s", from, m)
		if m == "hodor" {
			text = fmt.Sprintf("🚪🛑 Einlassstop\n\n%s", from)
		}
		message, _ := h.bot.SendGroupMessage(channel, text)
		e.Record.Set("tg_msg_id", message.MessageID)
		h.pb.Dao().SaveRecord(e.Record)
		h.AddAlert(NewAlertFromRecord(e.Record))
		return nil
	}
}

func (h *Handler) GetAfterUpdatedAlertHook() func(e *core.RecordUpdateEvent) error {
	return func(e *core.RecordUpdateEvent) error {
		log.Println("Updated alert:", e.Record)
		channel := e.Record.GetString("channel")
		msg := e.Record.GetString("msg")
		active := e.Record.GetBool("active")
		if msg == "hodor" {
			if e.Record.GetBool("active") {
				h.bot.SendGroupMessage(channel, "🚪🛑 Einlassstop")
			} else {
				h.bot.SendGroupMessage(channel, "🚪🟢 Tür wieder auf")
			}
			return nil
		}
		if active {
			return nil
		}
		from := strings.ToUpper(e.Record.GetString("from"))
		text := fmt.Sprintf("✅ %s ✅\n\n Hat sich erledigt", from)
		h.bot.EditGroupMessage(channel, e.Record.GetInt("tg_msg_id"), text)
		h.RemoveAlert(e.Record.Id)
		return nil
	}
}

func (h *Handler) SetToSeen(a *Alert) error {
	rec, err := h.pb.Dao().FindRecordById("alerts", a.ID)
	if err != nil {
		return err
	}
	rec.Set("seen", true)
	channel := rec.GetString("channel")
	from := strings.ToUpper(rec.GetString("from"))
	text := fmt.Sprintf("💬 %s 💬\n\n hierauf wurde geantwortet", from)
	h.bot.EditGroupMessage(channel, rec.GetInt("tg_msg_id"), text)

	return h.pb.Dao().SaveRecord(rec)
}
