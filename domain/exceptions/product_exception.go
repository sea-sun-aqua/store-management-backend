package exceptions
import (
	"errors"
)

var (
	ErrDuplicatedNameProduct = errors.New("duplicated product_name")
	ErrDuplicatedIDProduct = errors.New("duplicated product_id")
)