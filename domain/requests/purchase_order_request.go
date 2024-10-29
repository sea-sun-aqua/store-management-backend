package requests

type PurchaseOrderProductRequest struct {
	ProductID 	string `json:"product_id"`
	Amount 		int64 `json:"product_amount"`
}

type PurchaseOrderCreateRequest struct {
	PurchaseOrderID 	string `json:"purchase_order_id"`
	SupplierName 		string `json:"supplier_name"`
	PhoneNumber     	string `json:"supplier_phone_number"`
	Status 				string `json:"purchase_order_status"`
	CreatedDate     	string `json:"purchase_order_created_date"`
	Method 				string `json:"purchase_order_payment_method"`
	StaffID 			string `json:"staff_id"`
	Products 			[]PurchaseOrderProductRequest `json:"products"`
}

type PurchaseOrderUpdateStatusRequest struct {
	Status 		string `json:"purchase_order_status"`
}
