package requests

type ProductRegisterRequest struct {
	ProductID               string `json:"product_id"`
	Name 					string `json:"product_name"`
	PriceOfUnit 			float64 `json:"product_price"`
	Amount 					int64 `json:"product_amount"`
	QuantityOfSafetyStock 	int64 `json:"safety_stock_amount"`
}

type ProductUpdateAmountRequest struct {
	Amount 					int64 `json:"product_amount"`
}