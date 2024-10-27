package models

type SupplyOrderDetail struct {
	SupplyOrderID 		string `json:"supplyOrder" db:"supply_order_id"`
	ProductID 			string `json:"product_id" db:"product_id"`
	Amount 				int64 `json:"amount" db:"amount"`
}