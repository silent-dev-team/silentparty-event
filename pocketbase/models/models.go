package models

import "github.com/pocketbase/pocketbase/core"

type Model interface {
	Ticket | Hp | Marque | Dj | ShopItem
}

type Ticket struct {
	Id          string `json:"id" db:"id"`
	FirstName   string `json:"firstName" db:"firstName"`
	LastName    string `json:"lastName" db:"lastName"`
	Street      string `json:"street" db:"street"`
	HouseNumber string `json:"houseNumber" db:"houseNumber"`
	ZipCode     string `json:"zipCode" db:"zipCode"`
	Place       string `json:"place" db:"place"`
	Email       string `json:"email" db:"email"`
	Vvk         bool   `json:"vvk" db:"vvk"`
	Sold        bool   `json:"sold" db:"sold"`
	Filled      bool   `json:"filled" db:"filled"`
	Used        bool   `json:"used" db:"used"`
}

type Hp struct {
	Id     string `json:"id" db:"id"`
	Qr     string `json:"qr" db:"qr"`
	Lent   bool   `json:"lent" db:"lent"`
	Defect bool   `json:"defect" db:"defect"`
}

type Marque struct {
	Id   string `json:"id" db:"id"`
	Text string `json:"text" db:"text"`
}

type Dj struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Music     string `json:"music" db:"music"`
	Instagram string `json:"instagram" db:"instagram"`
	Color     string `json:"color" db:"color"`
}

type ShopItem struct {
	Id          string  `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float32 `json:"price" db:"price"`
	Img         string  `json:"img" db:"img"`
	Tags        string  `json:"tags" db:"tags"`
	Pfand       bool    `json:"pfand" db:"pfand"`
	PfandItem   string  `json:"pfand_item" db:"pfand_item"`
}

func GetAllTickets(app core.App) ([]Ticket, error) {
	var tickets []Ticket
	err := app.Dao().DB().
		NewQuery("SELECT * FROM tickets").
		All(&tickets)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func GetAllHps(app core.App) ([]Hp, error) {
	var hps []Hp
	err := app.Dao().DB().
		NewQuery("SELECT * FROM hp").
		All(&hps)
	if err != nil {
		return nil, err
	}
	return hps, nil
}

func GetAllMarques(app core.App) ([]Marque, error) {
	var marques []Marque
	err := app.Dao().DB().
		NewQuery("SELECT * FROM marquee").
		All(&marques)
	if err != nil {
		return nil, err
	}
	return marques, nil
}

func GetAllDjs(app core.App) ([]Dj, error) {
	var djs []Dj
	err := app.Dao().DB().
		NewQuery("SELECT * FROM djs").
		All(&djs)
	if err != nil {
		return nil, err
	}
	return djs, nil
}

func GetAllShopItems(app core.App) ([]ShopItem, error) {
	var items []ShopItem
	err := app.Dao().DB().
		NewQuery("SELECT * FROM shop_items").
		All(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
