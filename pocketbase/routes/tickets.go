package routes

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

type Transaction struct {
	Positions []Position `json:"positions"`
}

type Position struct {
	ItemId string `json:"itemId"`
	Amount int    `json:"amount"`
}

func GetTicketExists(p Path) ServerEventFunc {
	return func(e *core.ServeEvent) error {
		e.Router.GET(p, func(c echo.Context) error {
			id := c.PathParam("id")
			ticket, err := e.App.Dao().FindRecordById("tickets", id)
			if err != nil {
				return err
			}
			sold := ticket.Get("sold")
			if sold == nil || !sold.(bool) {
				c.JSON(404, map[string]string{"error": "ticket not sold"})
				return err
			}
			c.JSON(204, nil)
			return nil
		})
		return nil
	}
}
