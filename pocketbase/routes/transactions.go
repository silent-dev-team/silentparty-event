package routes

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func PostNewTransaction(p Path) ServerEventFunc {
	return func(e *core.ServeEvent) error {
		e.Router.POST("/api/v1/new-transaction", func(c echo.Context) error {
			positionsCollection, err1 := e.App.Dao().FindCollectionByNameOrId("positions")
			transactionsCollection, err2 := e.App.Dao().FindCollectionByNameOrId("transactions")

			if err1 != nil || err2 != nil {
				if err1 != nil {
					return err1
				}
				return err2
			}

			t := new(Transaction)
			err := c.Bind(&t)
			if err != nil {
				return err
			}
			posIds := make([]string, len(t.Positions))
			total := 0.0
			for _, p := range t.Positions {
				item, err := e.App.Dao().FindRecordById("shop_items", p.ItemId)
				if err != nil {
					return err
				}
				pos := models.NewRecord(positionsCollection)
				pos.RefreshId()
				pos.Set("item", p.ItemId)
				pos.Set("amount", p.Amount)
				pos.Set("price", item.GetFloat("price"))
				total += item.GetFloat("price") * float64(p.Amount)
				err = e.App.Dao().SaveRecord(pos)
				if err != nil {
					return err
				}
				posIds = append(posIds, pos.Id)
			}
			tr := models.NewRecord(transactionsCollection)
			tr.RefreshId()
			tr.Set("positions", posIds)
			tr.Set("total", total)
			err = e.App.Dao().SaveRecord(tr)
			if err != nil {
				return err
			}
			c.JSON(200, tr)
			return nil
		})
		return nil
	}
}
