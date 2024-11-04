package repositories

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
)

type ProductRepository interface {
	Create(ctx context.Context, req *requests.ProductRegisterRequest) error
	FindByName(ctx context.Context, name string) (*models.Product, error)
	FindByID(ctx context.Context, id string) (*models.Product, error)
	GetAll(ctx context.Context) ([]models.Product, error)
	UpdateAmountByID(ctx context.Context, id string, req *requests.ProductUpdateAmountRequest) error
}