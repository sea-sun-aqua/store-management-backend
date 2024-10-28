package requests

type OrderDetailCreateRequest struct {
	OrderID 	string `json:"order_id"`
	ProductID 	string `json:"product_id"`
	Amount 		int64 `json:"order_amount"`
}