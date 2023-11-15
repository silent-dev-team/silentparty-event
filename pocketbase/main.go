package main

import (
	"embed"
	"fmt"
	"log"
	"strings"

	"github.com/silent-dev-team/silentparty-event/pocketbase/alerts"
	"github.com/silent-dev-team/silentparty-event/pocketbase/hooks"
	_ "github.com/silent-dev-team/silentparty-event/pocketbase/migrations"
	"github.com/silent-dev-team/silentparty-event/pocketbase/routes"
	"github.com/silent-dev-team/silentparty-event/pocketbase/sse"
	"github.com/silent-dev-team/silentparty-event/pocketbase/tg"
	"github.com/silent-dev-team/silentparty-event/pocketbase/utils"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

//go:embed all:public
var public embed.FS

func main() {
	app := pocketbase.New()
	broker := sse.NewBroker(app)

	fmt.Println("Starting app")

	/* BOT */
	token := utils.Getenv("BOT_TOKEN")
	if token != "" {
		log.Println("Starting bot")
		groups := tg.GroupMap{
			"default":   utils.GetenvInt64("BOT_DEFAULT_GROUP"),
			"awareness": utils.GetenvInt64("BOT_AWARENESS_GROUP"),
		}
		bot, err := tg.NewBot(token, groups)
		if err != nil {
			log.Panicln("ERROR:", err)
		} else {
			go bot.MessageListener()
			alertHandler := alerts.NewHandler(app, bot)
			go alertHandler.StartReplyListener()
			// send message on alert
			app.OnRecordAfterCreateRequest("alerts").Add(alertHandler.GetAfterCreatedAlertHook())
			// send message on active off alert
			app.OnRecordAfterUpdateRequest("alerts").Add(alertHandler.GetAfterUpdatedAlertHook())
		}
	} else {
		log.Println("No bot token found")
	}

	/* CUSTOM ENDPOINTS */

	app.OnBeforeServe().Add(routes.GetTicketExists("/api/collections/tickets/exists/:id"))
	app.OnBeforeServe().Add(routes.PostNewTransaction("/api/v1/new-transaction"))
	app.OnBeforeServe().Add(routes.PostUnlink("/api/collections/ticket_hp/unlink"))
	app.OnBeforeServe().Add(routes.GetUserStats("/api/v1/userstats"))

	app.OnRealtimeAfterSubscribeRequest().Add(sse.OnSubscription(app))

	app.OnRealtimeBeforeMessageSend().Add(sse.OnMessage(app))

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
	app.OnRecordBeforeCreateRequest("tickets").
		Add(hooks.GeneratePin())

	// get position price from item
	app.OnRecordBeforeCreateRequest("positions").
		Add(hooks.PriceOfPosition(app))

	// calculate total price of transaction
	app.OnRecordBeforeCreateRequest("transactions").
		Add(hooks.TotalPriceOfTransaction(app))

	// dont allow illegal link
	app.OnRecordBeforeCreateRequest("ticket_hp").Add(hooks.DontAllowIllegalLink(app))

	// set hp to lent if it is linked to a ticket
	app.OnRecordAfterCreateRequest("ticket_hp").Add(hooks.SetHpToLentAndSetTicketToUsed(app))

	// trigger userstats on relevant changes -> implcitly send message to all clients that are subscribed to userstats
	app.OnRecordAfterUpdateRequest("tickets").Add(func(e *core.RecordUpdateEvent) error {
		broker.Send("userstats", []byte{})
		return nil
	})

	app.OnRecordAfterCreateRequest("ticket_hp").Add(func(e *core.RecordCreateEvent) error {
		broker.Send("userstats", []byte{})
		return nil
	})

	app.OnRecordAfterUpdateRequest("ticket_hp").Add(func(e *core.RecordUpdateEvent) error {
		broker.Send("userstats", []byte{})
		return nil
	})

	app.OnRecordAfterUpdateRequest("hp").Add(func(e *core.RecordUpdateEvent) error {
		broker.Send("userstats", []byte{})
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
