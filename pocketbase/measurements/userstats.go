package measurements

import (
	"log"

	"github.com/pocketbase/pocketbase/core"
	"github.com/silent-dev-team/silentparty-event/pocketbase/models"
)

type UserStats struct {
	Sells     int `json:"sells"`     // used all tickets
	Checked   int `json:"checked"`   // used all tickets
	UsedVvk   int `json:"usedVvk"`   // used vvk tickets
	UnusedVvk int `json:"unusedVvk"` // open vvk tickets
	Returned  int `json:"returned"`  // 0
	Current   int `json:"current"`   //
}

func GetUserStats(app core.App) (*UserStats, error) {
	tickets, err := models.GetAllTickets(app)
	if err != nil {
		log.Panicln(err)
	}

	var usedTickets int
	var usedVvkTickets int
	var unusedVvkTickets int
	for _, t := range tickets {
		if t.Used {
			usedTickets++
		}
		if t.Vvk && t.Sold && t.Used {
			usedVvkTickets++
		}
		if t.Vvk && t.Sold && !t.Used {
			unusedVvkTickets++
		}
	}

	var lentHps []models.Hp
	app.Dao().DB().
		NewQuery("SELECT * FROM hp WHERE lent=true").
		All(&lentHps)

	stats := &UserStats{
		Sells:     usedTickets,
		UsedVvk:   usedVvkTickets,
		UnusedVvk: unusedVvkTickets,
		Checked:   len(tickets),
		Returned:  0,
		Current:   len(lentHps),
	}

	return stats, nil
}
