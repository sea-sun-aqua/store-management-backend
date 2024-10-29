package repositories

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
)

type PurchaseOrderDetailRepository interface {
	GetAll(ctx context.Context) ([]models.PurchaseOrderDetail, error)
}