package main

type Transaction struct {
	Positions []Position `json:"positions"`
}

type Position struct {
	ItemId string `json:"itemId"`
	Amount int    `json:"amount"`
}

type Unlink struct {
	QR string `json:"qr"`
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

type Userstats struct {
	Sells    int `json:"sells"`    // used all tickets
	Checked  int `json:"checked"`  // used vvk tickets
	UsedVvk  int `json:"usedVvk"`  // used tickets (link Sells)
	Returned int `json:"returned"` // 0
	Current  int `json:"current"`  //
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
	Id          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	Img         string `json:"img" db:"img"`
	Tags        string `json:"tags" db:"tags"`
	Pfand       bool   `json:"pfand" db:"pfand"`
	PfandItem   string `json:"pfand_item" db:"pfand_item"`
}
