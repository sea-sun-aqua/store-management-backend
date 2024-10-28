package services

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/exceptions"
	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
)

type productService struct {
	productRepo repositories.ProductRepository
}


func NewProductService(productRepo repositories.ProductRepository) usecases.ProductUseCase {
	return &productService{
		productRepo: productRepo,
	}
}

func (p *productService) FindByName(ctx context.Context, name string) (*models.Product, error) {
	panic("unimplemented")
}

// FindByID implements usecases.ProductUseCase.
func (p *productService) FindByID(ctx context.Context, id string) (*models.Product, error) {
	panic("unimplemented")
}

func (p *productService) GetAllProducts(ctx context.Context) []models.Product {
	panic("unimplemented")
}

func (p *productService) Register(ctx context.Context, req *requests.ProductRegisterRequest) error {
	result, err := p.productRepo.FindByName(ctx, req.Name)

	if err != nil {
		return err
	}

	if result != nil {
		return exceptions.ErrDuplicatedNameProduct
	}

	result, err = p.productRepo.FindByID(ctx, req.ProductID)
	if err != nil {
		return err
	}

	if result != nil {
		return exceptions.ErrDuplicatedIDProduct
	}

	return p.productRepo.Create(ctx, req)
}
