package routes

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

type Unlink struct {
	QR string `json:"qr"`
}

func PostUnlink(p Path) ServerEventFunc {
	return func(e *core.ServeEvent) error {
		e.Router.POST(p, func(c echo.Context) error {
			unlink := new(Unlink)
			err := c.Bind(&unlink)
			if err != nil {
				c.JSON(400, map[string]string{"error": "bad request"})
				return err
			}
			hp, err := e.App.Dao().FindFirstRecordByFilter("hp", "qr={:qr}", dbx.Params{"qr": unlink.QR})
			if err != nil {
				c.JSON(404, map[string]string{"error": "headphone not found"})
				return err
			}
			ticketHp, err := e.App.Dao().FindFirstRecordByFilter("ticket_hp", "hp={:hp} && end = null", dbx.Params{"hp": hp.Id})
			if err != nil {
				c.JSON(404, map[string]string{"error": "relation not found"})
				return err
			}
			ticketHp.Set("end", types.NowDateTime())
			err = e.App.Dao().SaveRecord(ticketHp)

			if err != nil {
				c.JSON(500, map[string]string{"error": "internal server error - could not save end time"})
				return err
			}

			hp.Set("lent", false)
			err = e.App.Dao().SaveRecord(hp)

			if err != nil {
				c.JSON(500, map[string]string{"error": "internal server error - could not undlent hp"})
				return err
			}

			c.JSON(200, ticketHp)
			return nil
		})
		return nil
	}
}
