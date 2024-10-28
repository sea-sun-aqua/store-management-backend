package usecases

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
)

type ProductUseCase interface {
	Register(ctx context.Context, req *requests.ProductRegisterRequest) error
	GetAllProducts(ctx context.Context) []models.Product
	FindByName(ctx context.Context, name string) (*models.Product, error)
	FindByID(ctx context.Context, id string) (*models.Product, error)
}