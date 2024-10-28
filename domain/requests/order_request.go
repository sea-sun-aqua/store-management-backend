package requests


type OrderProductRequest struct {
	ProductID 	string `json:"product_id"`
	Amount 		string `json:"product_amount"`
}

type OrderCreateRequest struct {
	OrderID      string         `json:"order_id"`
	CustomerName string         `json:"customer_name"`
	Email        string         `json:"customer_email"`
	PhoneNumber  string         `json:"phone_number"`
	Address      string         `json:"customer_address"`
	Method       string         `json:"order_payment_method"`
	Status       string         `json:"order_status"`
	CreatedAt    string         `json:"order_created_date"`
	StaffID    	 string         `json:"staff_id"`
	Products   []OrderProductRequest `json:"products"`
}

type OrderUpdateStatusRequest struct {
	Status 		string `json:"order_status"`
}


