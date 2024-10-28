package repositories

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
)

type OrderRepository interface {
	Create(ctx context.Context, req *requests.OrderCreateRequest) error
	GetAll(ctx context.Context) ([]models.Order, error)
	FindByID(ctx context.Context, id string) (*models.Order, error)
	UpdateStatusByID(ctx context.Context, id string, req *requests.OrderUpdateStatusRequest) error
}