package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(
		func(db dbx.Builder) error {
			dao := daos.New(db)

			collection, err := dao.FindCollectionByNameOrId("shop_items")
			if err != nil {
				return err
			}

			rec, _ := dao.FindFirstRecordByFilter("shop_items", "title = 'AK Ticket'")
			if rec != nil {
				return nil
			}

			record := models.NewRecord(collection)
			record.Set("title", "AK Ticket")
			record.Set("description", "Ticket für die Abendkasse")
			record.Set("price", 10.0)
			record.Set("tags", []string{"ak", "ticket"})

			return dao.SaveRecord(record)
		}, nil)
}
