package model

type Expense struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Amount   int    `json:"amount"`
	Note     string `json:"note"`
}
