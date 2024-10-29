package models

type PurchaseOrderDetail struct {
	PurchaseOrderID 	string `json:"purchase_order_id" db:"purchase_order_id"`
	ProductID 			string 	`json:"product_id" db:"product_id"`
    Amount 				int64 `json:"purchase_order_amount" db:"purchase_order_amount"`
}