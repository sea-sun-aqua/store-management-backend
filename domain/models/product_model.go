package models

type Product struct {
	ProductID 				string 	`json:"product_id" db:"product_id"`
	Name 					string 	`json:"product_name" db:"product_name"`
	Price 					float64 `json:"product_price" db:"product_price"`
	Amount 					int64 	`json:"product_amount" db:"product_amount"`
	QuantityOfSafetyStock 	int64 	`json:"safety_stock_amount" db:"safety_stock_amount"`
}