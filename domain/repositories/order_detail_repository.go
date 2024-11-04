package repositories

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
)

type OrderDetailRepository interface {
	FindByOrderID(ctx context.Context, id string) ([]models.OrderDetail, error)
	FindByProductID(ctx context.Context, id string) ([]models.OrderDetail, error)
	GetAll(ctx context.Context) ([]models.OrderDetail, error)
}