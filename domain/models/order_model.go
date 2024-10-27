package models

const (
	OrderStatusTranferring  = "Tranferring"
	OrderStatusCompleted 	= "Completed"
)

type Order struct {
	OrderID        		string `json:"order_id" db:"order_id"`
	CustomerName      	string `json:"customer_name" db:"customer_name"`
	Email     			string `json:"email" db:"email"`
	PhoneNumber 		string `json:"phone_number" db:"phone_number"`
	Address 			string `json:"address" db:"address"`
	Method 				string `json:"method" db:"method"`
	Status 				string `json:"status" db:"status"`
	
}