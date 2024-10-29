package exceptions

import "errors"

var (
	ErrDuplicatedPurchaseIDOrder = errors.New("duplicated purchase_order_id")
	ErrPurchaseOrderIDNotFound = errors.New("purchase_order_id not found")
)