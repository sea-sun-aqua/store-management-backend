package models

type Product struct {
	ProductID 			string `json:"product_id" db:"product_id"`
	Name 				string `json:"name" db:"name"`
	Price 				float64 `json:"price" db:"price"`
	Amount 				int64 `json:"amount" db:"amount"`
	SafetyStockAmount 	int64 `json:"safety_stock_amount" db:"safety_stock_amount"`
	SupplyID 			string `json:"supply_id" db:"supply_id"`
}