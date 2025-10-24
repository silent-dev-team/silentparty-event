package measurements

import (
	"fmt"
	"log"

	"github.com/pocketbase/pocketbase/core"
	"github.com/silent-dev-team/silentparty-event/pocketbase/models"
)

type UserStats struct {
	UsedTickets int `json:"used_tickets"` // used all tickets
	UsedVVK     int `json:"used_vvk"`     // used vvk tickets
	UnusedVVK   int `json:"unused_vvk"`   // open vvk tickets
	UsedAK      int `json:"used_ak"`      // used ak tickets
	ActiveAK    int `json:"active_ak"`    // actived ak tickets
	ReservedAK  int `json:"reserved_ak"`  // ak tickets on hold
	OpenAK      int `json:"open_ak"`      // ak tickets open for sale
	LentHPs     int `json:"lent_hp"`      //
	UnusedHPs   int `json:"unused_hp"`    //
	Overbooks   int `json:"overbooks"`    //
}

func (us *UserStats) String() string {
	return "User Stats:\n" +
		fmt.Sprintf("- Used Tickets: %d\n", us.UsedTickets) +
		fmt.Sprintf("- Used VVK Tickets: %d\n", us.UsedVVK) +
		fmt.Sprintf("- Unused VVK Tickets: %d\n", us.UnusedVVK) +
		fmt.Sprintf("- Used AK Tickets: %d\n", us.UsedAK) +
		fmt.Sprintf("- Active AK Tickets: %d\n", us.ActiveAK) +
		fmt.Sprintf("- Open AK Tickets: %d\n", us.OpenAK) +
		fmt.Sprintf("- Reserved AK Tickets: %d\n", us.ReservedAK) +
		fmt.Sprintf("- Lent HPs: %d\n", us.LentHPs) +
		fmt.Sprintf("- Unused HPs: %d\n", us.UnusedHPs) +
		fmt.Sprintf("- Overbooks: %d\n", us.Overbooks)
}

type Number struct {
	Key   string `db:"key"`
	Value int    `db:"value"`
}

func GetUserStats(app core.App) (*UserStats, error) {
	tickets, err := models.GetAllTickets(app)
	if err != nil {
		log.Panicln(err)
	}
	hps, err := models.GetAllHps(app)
	if err != nil {
		log.Panicln(err)
	}

	var usedTickets int
	var usedVvkTickets int
	var unusedVvkTickets int
	var usedAkTickets int
	var activeAkTickets int
	var reservedAkTickets int
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
		if !t.Vvk && t.Sold && t.Used {
			usedAkTickets++
		}
		if !t.Vvk && t.Sold && !t.Used {
			activeAkTickets++
		}
		if !t.Vvk && !t.Sold && !t.Used && !t.Filled {
			reservedAkTickets++
		}
	}

	var defectHps []*models.Hp
	var lentHps []*models.Hp
	for _, hp := range hps {
		if hp.Defect {
			defectHps = append(defectHps, &hp)
			continue
		}
		if hp.Lent {
			lentHps = append(lentHps, &hp)
		}
	}
	unusedHps := len(hps) - len(lentHps) - len(defectHps)

	var overbooks Number
	err = app.DB().
		NewQuery("SELECT value FROM numbers n WHERE n.key = \"overbooks\" LIMIT 1").
		One(&overbooks)
	if err != nil {
		log.Println(err)
	}

	var teamReserve Number
	err = app.DB().
		NewQuery("SELECT value FROM numbers n WHERE n.key = \"team-reserve\" LIMIT 1").
		One(&teamReserve)
	if err != nil {
		log.Println(err)
	}

	clacOpenAKTickets := func() int {
		v := (unusedHps - unusedVvkTickets - activeAkTickets - teamReserve.Value) + min(
			overbooks.Value,
			unusedVvkTickets,
		)
		return max(v, 0)
	}

	stats := &UserStats{
		UsedTickets: usedTickets,
		UsedVVK:     usedVvkTickets,
		UnusedVVK:   unusedVvkTickets,
		UsedAK:      usedAkTickets,
		ActiveAK:    activeAkTickets,
		OpenAK:      clacOpenAKTickets(),
		ReservedAK:  reservedAkTickets,
		LentHPs:     len(lentHps) - teamReserve.Value,
		UnusedHPs:   unusedHps,
		Overbooks:   overbooks.Value,
	}

	return stats, nil
}
