package models

// User schema of the user table
type Stock struct {
	StockID int64  `json:"stockid"`
	Name    string `json:"name,omitempty"`
	Price   int    `json:"price,omitempty"`
	Company string `json:"company,omitempty"`
}
