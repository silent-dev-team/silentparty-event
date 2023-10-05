package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
					"id": "_pb_users_auth_",
					"name": "users",
					"type": "auth",
					"system": false,
					"schema": [
							{
									"system": false,
									"id": "users_name",
									"name": "name",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"system": false,
									"id": "users_avatar",
									"name": "avatar",
									"type": "file",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"maxSelect": 1,
											"maxSize": 5242880,
											"mimeTypes": [
													"image/jpeg",
													"image/png",
													"image/svg+xml",
													"image/gif",
													"image/webp"
											],
											"thumbs": null,
											"protected": false
									}
							}
					],
					"indexes": [],
					"listRule": "id = @request.auth.id",
					"viewRule": "id = @request.auth.id",
					"createRule": "",
					"updateRule": "id = @request.auth.id",
					"deleteRule": "id = @request.auth.id",
					"options": {
							"allowEmailAuth": true,
							"allowOAuth2Auth": true,
							"allowUsernameAuth": true,
							"exceptEmailDomains": null,
							"manageRule": null,
							"minPasswordLength": 8,
							"onlyEmailDomains": null,
							"requireEmail": false
					}
			},
			{
					"id": "fnt1n74ihcaofvy",
					"name": "tickets",
					"type": "base",
					"system": false,
					"schema": [
							{
									"system": false,
									"id": "s333zfqz",
									"name": "firstName",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"system": false,
									"id": "lplzdwo6",
									"name": "lastName",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"system": false,
									"id": "xyeebfdq",
									"name": "street",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"system": false,
									"id": "ngmkp7br",
									"name": "housenumber",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"system": false,
									"id": "khwbqnbd",
									"name": "zipCode",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"system": false,
									"id": "so4dikq2",
									"name": "place",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"system": false,
									"id": "f4hqk1fk",
									"name": "email",
									"type": "email",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"exceptDomains": null,
											"onlyDomains": null
									}
							},
							{
									"system": false,
									"id": "9rlcgjaz",
									"name": "sold",
									"type": "bool",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {}
							},
							{
									"system": false,
									"id": "s5qooewb",
									"name": "filled",
									"type": "bool",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {}
							},
							{
									"system": false,
									"id": "dkfgw3lf",
									"name": "validated",
									"type": "bool",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {}
							},
							{
									"system": false,
									"id": "opytsfoy",
									"name": "used",
									"type": "bool",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {}
							},
							{
									"system": false,
									"id": "qoicngue",
									"name": "_pin",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": 4,
											"max": 4,
											"pattern": "^[0-9]{4}$"
									}
							}
					],
					"indexes": [],
					"listRule": null,
					"viewRule": "((filled = true && @request.query.pin = _pin) || filled = false) && sold = true",
					"createRule": null,
					"updateRule": "used = false",
					"deleteRule": null,
					"options": {}
			},
			{
					"id": "nhw1dv5dnt6yxeh",
					"name": "shop_items",
					"type": "base",
					"system": false,
					"schema": [
							{
									"system": false,
									"id": "vibp24pt",
									"name": "title",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"system": false,
									"id": "hszg1dyl",
									"name": "description",
									"type": "text",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"pattern": ""
									}
							},
							{
									"system": false,
									"id": "fwqvx2jo",
									"name": "price",
									"type": "number",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"noDecimal": false
									}
							}
					],
					"indexes": [
							"CREATE UNIQUE INDEX 'idx_xdWr4LF' ON 'shop_items' ('title')"
					],
					"listRule": null,
					"viewRule": null,
					"createRule": null,
					"updateRule": null,
					"deleteRule": null,
					"options": {}
			},
			{
					"id": "o5kg9huo5umgqun",
					"name": "transactions",
					"type": "base",
					"system": false,
					"schema": [
							{
									"system": false,
									"id": "vizjb3yv",
									"name": "positions",
									"type": "relation",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"collectionId": "3ocopq5t68o1565",
											"cascadeDelete": false,
											"minSelect": null,
											"maxSelect": null,
											"displayFields": null
									}
							},
							{
									"system": false,
									"id": "3bit5x8c",
									"name": "total",
									"type": "number",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"noDecimal": false
									}
							},
							{
									"system": false,
									"id": "hoy7nwkt",
									"name": "storno",
									"type": "bool",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {}
							}
					],
					"indexes": [],
					"listRule": null,
					"viewRule": null,
					"createRule": null,
					"updateRule": null,
					"deleteRule": null,
					"options": {}
			},
			{
					"id": "3ocopq5t68o1565",
					"name": "positions",
					"type": "base",
					"system": false,
					"schema": [
							{
									"system": false,
									"id": "0yvanlyc",
									"name": "item",
									"type": "relation",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"collectionId": "nhw1dv5dnt6yxeh",
											"cascadeDelete": false,
											"minSelect": null,
											"maxSelect": 1,
											"displayFields": null
									}
							},
							{
									"system": false,
									"id": "qolqrrs6",
									"name": "price",
									"type": "number",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"noDecimal": false
									}
							},
							{
									"system": false,
									"id": "gpb1aycd",
									"name": "amount",
									"type": "number",
									"required": false,
									"presentable": false,
									"unique": false,
									"options": {
											"min": null,
											"max": null,
											"noDecimal": false
									}
							}
					],
					"indexes": [
							"CREATE INDEX 'idx_86HZizz' ON 'positions' ('item')"
					],
					"listRule": null,
					"viewRule": null,
					"createRule": null,
					"updateRule": null,
					"deleteRule": null,
					"options": {}
			}
	]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
