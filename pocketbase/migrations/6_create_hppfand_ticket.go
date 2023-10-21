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

			rec, _ := dao.FindFirstRecordByFilter("shop_items", "title = 'HP Pfand'")
			if rec != nil {
				return nil
			}

			record := models.NewRecord(collection)
			record.Set("title", "HP Pfand")
			record.Set("description", "Das ist der Pfand für die Kopfhörer")
			record.Set("price", -5)
			record.Set("tags", []string{"ak"})

			return dao.SaveRecord(record)

		}, nil)
}
