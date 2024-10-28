package exceptions

import "errors"

var (
	ErrDuplicatedIDOrder = errors.New("duplicated id_order")
	ErrStatusInvalid = errors.New("status_order_invalid")
	ErrOrderIDNotFound = errors.New("id_order_not_found")
)