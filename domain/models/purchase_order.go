package models

const (
	PurchaseOrderStatusTranferring = "Transferring"
	PurchaseOrderStatusCompleted   = "Completed"
	PurchaseOrderStatusCanceled    = "Canceled"
)

type PurchaseOrder struct {
	PurchaseOrderID 	string `json:"purchase_order_id" db:"purchase_order_id"`
	SupplierName 		string `json:"supplier_name" db:"supplier_name"`
	PhoneNumber     	string `json:"supplier_phone_number" db:"supplier_phone_number"`
	Status 				string `json:"purchase_order_status" db:"purchase_order_status"`
	CreatedDate     	string `json:"purchase_order_created_date" db:"purchase_order_created_date"`
	Method 				string `json:"purchase_order_payment_method" db:"purchase_order_payment_method"`
	StaffID 			string `json:"staff_id" db:"staff_id"`
}