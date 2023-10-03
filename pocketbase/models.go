package main

type Transaction struct {
	Positions []Position `json:"positions"`
}

type Position struct {
	ItemId string `json:"itemId"`
	Amount int    `json:"amount"`
}
