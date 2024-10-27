package models

const (
	SupplyOrderStatusTranferring = "TRANFERRING"
	SupplyOrderStatusPack = "PACK"
	SupplyOrderStatusDelivery = "DERIVERY"
	SupplyOrderStatusCompleted = "COMPLETED"
)

type SupplyOrder struct {
	SupplyOrderID 	string `json:"supplyOrder" db:"supply_order_id"`
	SupplierName 	string `json:"supplier_name" db:"supplier_name"`
	Status 			string `json:"status" db:"status"`
	CreatedDate     string `json:"created_date" db:"created_date"`
	PhoneNumber     string `json:"phone_number" db:"phone_number"`
	Method 			string `json:"method" db:"method"`
	StaffID 		string `json:"staff_id" db:"staff_id"`
}