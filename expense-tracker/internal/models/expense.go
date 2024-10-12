package models

type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"name"`
	Date        string  `json:"date"`
	Price       float64 `json:"price"`
}
