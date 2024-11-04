package repositories

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
)

type PurchaseOrderRepository interface {
	Create(ctx context.Context, req *requests.PurchaseOrderCreateRequest) error
	GetAll(ctx context.Context) ([]models.PurchaseOrder, error)
	FindByID(ctx context.Context, id string) (*models.PurchaseOrder, error)
	UpdateStatusByID(ctx context.Context, id string, req *requests.PurchaseOrderUpdateStatusRequest) error
}