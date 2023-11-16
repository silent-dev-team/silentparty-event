package hooks

import (
	"fmt"
	"net/mail"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/silent-dev-team/silentparty-event/pocketbase/utils"
)

func GeneratePin() func(e *core.RecordCreateEvent) error {
	return func(e *core.RecordCreateEvent) error {
		oldPin := e.Record.Get("_pin")
		fmt.Println("old pin: ", oldPin)
		if oldPin != "" {
			return nil
		}
		pin := utils.GeneratePin()
		e.Record.Set("_pin", pin)
		return nil
	}
}

func SendMail(app core.App) func(e *core.RecordUpdateEvent) error {
	return func(e *core.RecordUpdateEvent) error {
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
			Subject: "Dein personalisiertes Ticket für die Silent Party",
			HTML:    "<h1>Herzlichen Glückwunsch zur Personalisierung</h1><p>Hallo " + e.Record.GetString("firstName") + ",</p><p>du hast dein Ticket für die Silentparty am 18.11.23 erfolgreich personalisiert.</p><p>Wir freuen uns auf dich!</p><p>Dein Silent Party Team</p><p><img src=\"https://www.silentparty-hannover.de/images/logo.png\" alt=\"Silent Party Logo\" width=\"200\" height=\"200\"></p><p><a href=\"https://www.silentparty-hannover.de\">www.silentparty-hannover.de</a></p>",
			// bcc, cc, attachments and custom headers are also supported...
		}

		// send message
		err := app.NewMailClient().Send(message)
		if err != nil {
			fmt.Println(err)
		}
		return err
	}
}
