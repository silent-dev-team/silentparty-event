package measurements

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/silent-dev-team/silentparty-event/pocketbase/models"
)

type UserStats struct {
	Sells    int `json:"sells"`    // used all tickets
	Checked  int `json:"checked"`  // used vvk tickets
	UsedVvk  int `json:"usedVvk"`  // used tickets (link Sells)
	Returned int `json:"returned"` // 0
	Current  int `json:"current"`  //
}

func GetUserStats(app core.App) (*UserStats, error) {
	var usedTickets []models.Ticket
	app.Dao().DB().
		NewQuery("SELECT * FROM tickets WHERE used=true").
		All(&usedTickets)

	var usedVvkTickets int
	for _, t := range usedTickets {
		if t.Vvk {
			usedVvkTickets++
		}
	}

	var lentHps []models.Hp
	app.Dao().DB().
		NewQuery("SELECT * FROM hp WHERE lent=true").
		All(&lentHps)

	stats := &UserStats{
		Sells:    len(usedTickets),
		UsedVvk:  usedVvkTickets,
		Checked:  len(usedTickets),
		Returned: 0,
		Current:  len(lentHps),
	}

	return stats, nil
}
