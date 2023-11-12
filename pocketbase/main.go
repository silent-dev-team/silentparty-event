package main

import (
	"embed"
	"fmt"
	"log"
	"strings"

	_ "github.com/silent-dev-team/silentparty-event/pocketbase/migrations"
	"github.com/silent-dev-team/silentparty-event/pocketbase/tg"
	"github.com/silent-dev-team/silentparty-event/pocketbase/utils"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

//go:embed all:public
var public embed.FS

func main() {
	app := pocketbase.New()

	/* BOT */
	groups := tg.GroupMap{
		"default": utils.GetenvInt64("BOT_DEFAULT_GROUP"),
	}
	bot, err := tg.NewBot(utils.Getenv("BOT_TOKEN"), groups)
	if err != nil {
		log.Fatal(err)
	}
	go bot.Run()

	/* CUSTOM ENDPOINTS */

	// check if ticket exists
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/collections/tickets/exists/:id", func(c echo.Context) error {
			id := c.PathParam("id")
			ticket, err := app.Dao().FindRecordById("tickets", id)
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
			c.JSON(200, tr)
			return nil
		})
		return nil
	})

	// unlink hp from ticket
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/api/collections/ticket_hp/unlink", func(c echo.Context) error {
			unlink := new(Unlink)
			err := c.Bind(&unlink)
			if err != nil {
				c.JSON(400, map[string]string{"error": "bad request"})
				return err
			}
			hp, err := app.Dao().FindFirstRecordByFilter("hp", "qr={:qr}", dbx.Params{"qr": unlink.QR})
			if err != nil {
				c.JSON(404, map[string]string{"error": "headphone not found"})
				return err
			}
			ticketHp, err := app.Dao().FindFirstRecordByFilter("ticket_hp", "hp={:hp} && end = null", dbx.Params{"hp": hp.Id})
			if err != nil {
				c.JSON(404, map[string]string{"error": "relation not found"})
				return err
			}
			ticketHp.Set("end", types.NowDateTime())
			err = app.Dao().SaveRecord(ticketHp)

			if err != nil {
				c.JSON(500, map[string]string{"error": "internal server error - could not save end time"})
				return err
			}

			hp.Set("lent", false)
			err = app.Dao().SaveRecord(hp)

			if err != nil {
				c.JSON(500, map[string]string{"error": "internal server error - could not undlent hp"})
				return err
			}

			c.JSON(200, ticketHp)
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
		oldPin := e.Record.Get("_pin")
		fmt.Println("old pin: ", oldPin)
		if oldPin != "" {
			return nil
		}
		pin := utils.GeneratePin()
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

	app.OnRecordBeforeCreateRequest("ticket_hp").Add(func(e *core.RecordCreateEvent) error {
		app.Dao().ExpandRecord(e.Record, []string{"hp", "ticket"}, nil)
		hp := e.Record.ExpandedOne("hp")
		ticket := e.Record.ExpandedOne("ticket")
		if hp.GetBool("lent") {
			return fmt.Errorf("headphone is already lent")
		}
		if !ticket.GetBool("sold") {
			return fmt.Errorf("ticket is not sold")
		}
		if ticket.GetBool("used") {
			return fmt.Errorf("ticket is already used")
		}
		//TODO: if any link is open, dont allow a new one
		return nil
	})

	// set hp to lent if it is linked to a ticket
	app.OnRecordAfterCreateRequest("ticket_hp").Add(func(e *core.RecordCreateEvent) error {
		hpId := e.Record.GetString("hp")
		ticketId := e.Record.GetString("ticket")
		if hpId == "" || ticketId == "" {
			return nil
		}
		e.Record.Set("start", types.NowDateTime())
		app.Dao().SaveRecord(e.Record)

		hp, err := app.Dao().FindRecordById("hp", hpId)
		if err != nil {
			return err
		}
		hp.Set("lent", true)
		err = app.Dao().SaveRecord(hp)
		if err != nil {
			return err
		}

		ticket, err := app.Dao().FindRecordById("tickets", ticketId)
		if err != nil {
			return err
		}
		ticket.Set("used", true)
		err = app.Dao().SaveRecord(ticket)
		if err != nil {
			return err
		}

		return nil
	})

	// set hp to not lent if it is unlinked from a ticket
	app.OnRecordAfterUpdateRequest("ticket_hp").Add(func(e *core.RecordUpdateEvent) error {
		endTime := e.Record.GetDateTime("end")
		if endTime.IsZero() {
			return nil
		}
		hpId := e.Record.GetString("hp")
		ticketId := e.Record.GetString("ticket")
		if hpId == "" || ticketId == "" {
			return nil
		}
		hp, err := app.Dao().FindRecordById("hp", hpId)
		if err != nil {
			return err
		}
		hp.Set("lent", false)
		err = app.Dao().SaveRecord(hp)
		if err != nil {
			return err
		}
		return nil
	})

	// send message on alert
	app.OnRecordAfterCreateRequest("alerts").Add(func(e *core.RecordCreateEvent) error {
		m := e.Record.GetString("msg")
		from := strings.ToUpper(e.Record.GetString("from"))
		text := fmt.Sprintf("⚠️ %s ⚠️\n\n%s", from, m)
		message, _ := bot.SendGroupMessage("default", text)
		e.Record.Set("tg_msg_id", message.MessageID)
		app.Dao().SaveRecord(e.Record)
		return nil
	})

	// send message on active off alert
	app.OnRecordAfterUpdateRequest("alerts").Add(func(e *core.RecordUpdateEvent) error {
		from := strings.ToUpper(e.Record.GetString("from"))
		active := e.Record.GetBool("active")
		if active {
			seen := e.Record.GetBool("seen")
			if seen {
				text := fmt.Sprintf("🚶‍♂️ %s 🚶‍♂️\n\n Jemand ist auf dem Weg", from)
				bot.EditGroupMessage(e.Record.GetInt("tg_msg_id"), text, "default")
			}
			return nil
		}
		text := fmt.Sprintf("✅ %s ✅\n\n Hat sich erledigt", from)
		bot.EditGroupMessage(e.Record.GetInt("tg_msg_id"), text, "default")
		return nil
	})

	// send mail after ticket is sold
	// app.OnRecordAfterUpdateRequest("tickets").Add(func(e *core.RecordUpdateEvent) error {
	// 	sold := e.Record.GetBool("sold")
	// 	filled := e.Record.GetBool("filled")
	// 	used := e.Record.GetBool("used")
	// 	if used || !sold || (sold && !filled) {
	// 		return nil
	// 	}
	// 	message := &mailer.Message{
	// 		From: mail.Address{
	// 			Address: app.Settings().Meta.SenderAddress,
	// 			Name:    app.Settings().Meta.SenderName,
	// 		},
	// 		To:      []mail.Address{{Address: e.Record.GetString("email")}},
	// 		Subject: "Dein personalisiertes Ticket für die Silent Party",
	// 		HTML:    "<h1>Herzlichen Glückwunsch zur Personalisierung</h1><p>Hallo " + e.Record.GetString("firstName") + ",</p><p>du hast dein Ticket für die Silentparty am 18.11.23 erfolgreich personalisiert.</p><p>Wir freuen uns auf dich!</p><p>Dein Silent Party Team</p><p><img src=\"https://www.silentparty-hannover.de/images/logo.png\" alt=\"Silent Party Logo\" width=\"200\" height=\"200\"></p><p><a href=\"https://www.silentparty-hannover.de\">www.silentparty-hannover.de</a></p>",
	// 		// bcc, cc, attachments and custom headers are also supported...
	// 	}

	// 	// send message
	// 	err := app.NewMailClient().Send(message)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	return err
	// })

	/* SERVE */

	// serve frontend

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(echo.MustSubFS(public, "public"), true))
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/assets/*", apis.StaticDirectoryHandler(echo.MustSubFS(public, "public/assets"), true))
		return nil
	})

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

func SetToSeen(app *pocketbase.PocketBase, alertID string) error {
	alert, err := app.Dao().FindRecordById("alerts", alertID)
	if err != nil {
		return err
	}
	alert.Set("seen", true)
	return app.Dao().SaveRecord(alert)
}
