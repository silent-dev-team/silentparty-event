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
	"github.com/pocketbase/pocketbase/tools/mailer"
)

//go:embed all:public
var public embed.FS

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(echo.MustSubFS(public, "public"), true))
		return nil
	})

	app.OnRecordViewRequest().Add(LowdashFilterRecordViewEvent)
	app.OnRecordBeforeCreateRequest("tickets").Add(func(e *core.RecordCreateEvent) error {
		e.Record.Set("_pin", utils.GeneratePin())
		return nil
	})

	app.OnRecordBeforeCreateRequest("tickets").Add(func(e *core.RecordCreateEvent) error {
		message := &mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			To:      []mail.Address{{Address: e.Record.Email()}},
			Subject: "Dein Ticket für die Silent Party",
			HTML:    "<h1>Dein Ticket für die Silent Party</h1><p>Hallo " + e.Record.GetString("firstName") + ",</p><p>hier ist dein Ticket für die Silent Party.</p><p>Bitte bringe es ausgedruckt oder auf deinem Smartphone mit.</p><p>Wir freuen uns auf dich!</p><p>Dein Silent Party Team</p><p><img src=\"https://www.silentparty-hannover.de/images/logo.png\" alt=\"Silent Party Logo\" width=\"200\" height=\"200\"></p><p><a href=\"https://www.silentparty-hannover.de\">www.silentparty-hannover.de</a></p>",
			// bcc, cc, attachments and custom headers are also supported...
		}

		err := app.NewMailClient().Send(message)
		if err != nil {
			fmt.Println(err)
		}
		return err
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}

func LowdashFilterRecordViewEvent(e *core.RecordViewEvent) error {
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
}
