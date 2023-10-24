package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	col := "shop_items"
	title := "Becherpfand"
	price := 5
	plusTitle := title + " +"
	minusTitle := title + " -"
	tags := []string{"bar"}

	m.Register(
		func(db dbx.Builder) error {
			dao := daos.New(db)

			collection, err := dao.FindCollectionByNameOrId(col)
			if err != nil {
				return err
			}

			rec1, _ := dao.FindFirstRecordByFilter(col, "title = '{:t}'", dbx.Params{"t": plusTitle})
			if rec1 == nil {
				record := models.NewRecord(collection)
				record.Set("title", plusTitle)
				record.Set("description", "Pfand einnehmen")
				record.Set("price", price)
				record.Set("tags", tags)

				dao.SaveRecord(record)
			}

			rec2, _ := dao.FindFirstRecordByFilter(col, "title = '{:t}'", dbx.Params{"t": minusTitle})
			if rec2 == nil {
				record := models.NewRecord(collection)
				record.Set("title", minusTitle)
				record.Set("description", "Pfand ausgeben")
				record.Set("price", price)
				record.Set("tags", tags)

				dao.SaveRecord(record)
			}

			return nil

		}, nil)
}
