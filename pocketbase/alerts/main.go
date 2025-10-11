package alerts

import (
	"log"

	"slices"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/silent-dev-team/silentparty-event/pocketbase/tg"
)

type Handler struct {
	pb     *pocketbase.PocketBase
	bot    *tg.Bot
	Alerts []*Alert
}

func NewHandler(pb *pocketbase.PocketBase, bot *tg.Bot) *Handler {
	return &Handler{
		pb:     pb,
		bot:    bot,
		Alerts: []*Alert{},
	}
}

func (h *Handler) AddAlert(a *Alert) {
	h.Alerts = append(h.Alerts, a)
}

func (h *Handler) RemoveAlert(id string) {
	for i, a := range h.Alerts {
		if a.ID == id {
			h.Alerts = slices.Delete(h.Alerts, i, i+1)
			break
		}
	}
}

// Problem: Updates werden nicht empfangen, da der Kanal an anderer Stelle schon gefangen wird
func (h *Handler) StartReplyListener() {
	for reply := range h.bot.Replys {
		log.Println(".....REPLY!!:", reply)
		for _, a := range h.Alerts {
			log.Println(a)
			if a.TgMsgID == reply.ReplyToMessage.MessageID {
				h.SetToSeen(a)
			}
		}
	}
}

type Alert struct {
	ID      string `json:"id"`
	From    string `json:"from"`
	Msg     string `json:"msg"`
	Channel string `json:"channel"`
	Active  bool   `json:"active"`
	Seen    bool   `json:"seen"`
	TgMsgID int    `json:"tg_msg_id"`
}

func NewAlertFromRecord(r *models.Record) *Alert {
	return &Alert{
		ID:      r.Id,
		From:    r.GetString("from"),
		Msg:     r.GetString("msg"),
		Active:  r.GetBool("active"),
		Seen:    r.GetBool("seen"),
		TgMsgID: r.GetInt("tg_msg_id"),
	}
}
