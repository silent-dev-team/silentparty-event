package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	col := "shop_items"
	title := "Becher"
	price := 5
	tags := []string{"bar"}

	m.Register(
		func(db dbx.Builder) error {
			dao := daos.New(db)

			collection, err := dao.FindCollectionByNameOrId(col)
			if err != nil {
				return err
			}

			rec, _ := dao.FindFirstRecordByFilter(col, "title = '{:t}'", dbx.Params{"t": title})
			if rec == nil {
				record := models.NewRecord(collection)
				record.Set("title", title)
				record.Set("description", "Ausgabe eines Bechers gegen Pfand")
				record.Set("price", price)
				record.Set("tags", tags)
				record.Set("pfand", true)

				dao.SaveRecord(record)
			}

			return nil

		}, nil)
}
