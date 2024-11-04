package models

const (
	OrderStatusTranferring   = "Transferring"
	OrderStatusPack			 = "Pack"
	OrderStatusDeliver 		 = "Delivery"
	OrderStatusCompleted 	 = "Completed"
	OrderStatusCanceled 	 = "Canceled"
	OrderStatusPending	 	 = "Pending"
)

type Order struct {
	OrderID        		string `json:"order_id" db:"order_id"`
	CustomerName      	string `json:"customer_name" db:"customer_name"`
	Email     			string `json:"customer_email" db:"customer_email"`
	PhoneNumber 		string `json:"phone_number" db:"phone_number"`
	Address 			string `json:"customer_address" db:"customer_address"`
	Method 				string `json:"order_payment_method" db:"order_payment_method"`
	Status 				string `json:"order_status" db:"order_status"`
	CreatedAt 			string `json:"order_created_date" db:"order_created_date"`
	StaffID 			string `json:"staff_id" db:"staff_id"`
}

