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

			record := models.NewRecord(collection)
			record.Set("title", "VVK Ticket")
			record.Set("description", "Ticket für den Vorverkauf")
			record.Set("price", 7.0)
			record.Set("tags", []string{"vvk", "ticket"})

			return dao.SaveRecord(record)
		}, nil)
}
