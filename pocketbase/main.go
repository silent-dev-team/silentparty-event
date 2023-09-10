package main

import (
	"embed"
	"log"
	"strings"

	_ "github.com/silent-dev-team/silentparty-event/pocketbase/migrations"
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

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(echo.MustSubFS(public, "public"), true))
		return nil
	})

	app.OnRecordViewRequest().Add(LowdashFilterRecordViewEvent)
	app.OnRecordBeforeCreateRequest("tickets").Add(func(e *core.RecordCreateEvent) error {
		e.Record.Set("_salt", utils.GenerateRandomString(15))
		return nil
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
