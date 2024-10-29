package usecases

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
)

type PurchaseOrderDetailUseCase interface {
	GetAll(ctx context.Context) ([]models.PurchaseOrderDetail, error)
}