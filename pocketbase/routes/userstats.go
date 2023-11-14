package routes

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/silent-dev-team/silentparty-event/pocketbase/measurements"
)

func GetUserStats(p Path) ServerEventFunc {
	return func(e *core.ServeEvent) error {
		e.Router.GET(p, func(c echo.Context) error {
			stats, _ := measurements.GetUserStats(e.App)
			c.JSON(200, stats)
			return nil
		})
		return nil
	}
}
