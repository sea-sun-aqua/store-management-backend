package models

type OrderDetail struct {
	OrderID        		string `json:"order_id" db:"order_id"`
	ProductID 			string `json:"product_id" db:"product_id"`
	Amount 				int64 `json:"amount" db:"amount"`
}