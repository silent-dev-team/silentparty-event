package main

import (
	"embed"
	"fmt"
	"log"
	"net/mail"
	"strings"

	_ "github.com/silent-dev-team/silentparty-event/pocketbase/migrations"
	"github.com/silent-dev-team/silentparty-event/pocketbase/utils"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

//go:embed all:public
var public embed.FS

func main() {
	app := pocketbase.New()

	/* CUSTOM ENDPOINTS */

	// serve frontend
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(echo.MustSubFS(public, "public"), true))
		return nil
	})

	// create new transaction
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/api/v1/new-transaction", func(c echo.Context) error {
			positionsCollection, err1 := app.Dao().FindCollectionByNameOrId("positions")
			transactionsCollection, err2 := app.Dao().FindCollectionByNameOrId("transactions")

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
				item, err := app.Dao().FindRecordById("shop_items", p.ItemId)
				if err != nil {
					return err
				}
				pos := models.NewRecord(positionsCollection)
				pos.RefreshId()
				pos.Set("item", p.ItemId)
				pos.Set("amount", p.Amount)
				pos.Set("price", item.GetFloat("price"))
				total += item.GetFloat("price") * float64(p.Amount)
				err = app.Dao().SaveRecord(pos)
				if err != nil {
					return err
				}
				posIds = append(posIds, pos.Id)
			}
			tr := models.NewRecord(transactionsCollection)
			tr.RefreshId()
			tr.Set("positions", posIds)
			tr.Set("total", total)
			err = app.Dao().SaveRecord(tr)
			if err != nil {
				return err
			}
			return nil
		})
		return nil
	})

	/* CUSTOM HOOKS */

	// filter fields that start with _
	app.OnRecordViewRequest().Add(func(e *core.RecordViewEvent) error {
		resp := make(map[string]interface{})

		for k, v := range e.Record.ColumnValueMap() {
			if !strings.HasPrefix(k, "_") {
				resp[k] = v
			}
		}

		err := e.HttpContext.JSON(200, resp)
		if err != nil {
			return err
		}

		return nil
	})

	// generate pin
	app.OnRecordBeforeCreateRequest("tickets").Add(func(e *core.RecordCreateEvent) error {
		pin := utils.GeneratePin()
		fmt.Println(pin)
		e.Record.Set("_pin", pin)
		return nil
	})

	// get position price from item
	app.OnRecordBeforeCreateRequest("positions").Add(func(e *core.RecordCreateEvent) error {
		app.Dao().ExpandRecord(e.Record, []string{"item"}, nil)
		item := e.Record.ExpandedOne("item")
		e.Record.Set("price", item.GetFloat("price"))
		return nil
	})

	// calculate total price of transaction
	app.OnRecordBeforeCreateRequest("transactions").Add(func(e *core.RecordCreateEvent) error {
		app.Dao().ExpandRecord(e.Record, []string{"positions"}, nil)
		positions := e.Record.ExpandedAll("positions")
		e.Record.Set("total", calculateTotal(positions))
		return nil
	})

	// send mail after ticket is sold
	app.OnRecordAfterUpdateRequest("tickets").Add(func(e *core.RecordUpdateEvent) error {
		sold := e.Record.GetBool("sold")
		filled := e.Record.GetBool("filled")
		used := e.Record.GetBool("used")
		if used || !sold || (sold && !filled) {
			return nil
		}
		message := &mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: e.Record.GetString("email")}},
			Subject: "Dein Ticket für die Silent Party",
			HTML:    "<h1>Dein Ticket für die Silent Party</h1><p>Hallo " + e.Record.GetString("firstName") + ",</p><p>hier ist dein Ticket für die Silent Party.</p><p>Bitte bringe es ausgedruckt oder auf deinem Smartphone mit.</p><p>Wir freuen uns auf dich!</p><p>Dein Silent Party Team</p><p><img src=\"https://www.silentparty-hannover.de/images/logo.png\" alt=\"Silent Party Logo\" width=\"200\" height=\"200\"></p><p><a href=\"https://www.silentparty-hannover.de\">www.silentparty-hannover.de</a></p>",
			// bcc, cc, attachments and custom headers are also supported...
		}

		// send message
		err := app.NewMailClient().Send(message)
		if err != nil {
			fmt.Println(err)
		}
		return err
	})

	/* SERVE */

	// serve app
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}

func calculateTotal(positions []*models.Record) float64 {
	total := 0.0
	for _, p := range positions {
		total += p.GetFloat("price") * float64(p.GetInt("amount"))
	}
	return total
}
