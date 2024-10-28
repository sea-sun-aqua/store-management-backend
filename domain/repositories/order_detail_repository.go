package repositories

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
)

type OrderDetailRepository interface {
	Create(ctx context.Context, req *requests.OrderDetailCreateRequest) error
	FindByOrderID(ctx context.Context, id string) ([]models.OrderDetail, error)
	FindByProductID(ctx context.Context, id string) ([]models.OrderDetail, error)
}