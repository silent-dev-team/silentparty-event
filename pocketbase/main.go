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
	app.OnBeforeServe().Add(routes.PostNewTransaction("/api/v1/new-transaction", broker))
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

	// send mail on ticket update
	// app.OnRecordAfterUpdateRequest("tickets").Add(hooks.SendMail(app))

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
