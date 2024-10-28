package requests

type ProductRegisterRequest struct {
	ProductID               string `json:"product_id"`
	Name 					string `json:"product_name"`
	PriceOfUnit 			float64 `json:"price_of_unit"`
	QuantityOfSafetyStock 	int64 `json:"quantity_of_safety_stock"`
	Amount 					int64 `json:"product_amount"`
}