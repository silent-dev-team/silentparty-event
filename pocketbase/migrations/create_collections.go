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
									"id": "miodo3vo",
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
					"createRule": "",
					"updateRule": "",
					"deleteRule": "",
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
