package usecases

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
)

type PurchaseOrderUseCase interface {
	Create(ctx context.Context, req *requests.PurchaseOrderCreateRequest) error
	UpdateStatusByID(ctx context.Context, id string, req *requests.PurchaseOrderUpdateStatusRequest) (*models.PurchaseOrder, error)
	GetAll(ctx context.Context) ([]models.PurchaseOrder, error)
}