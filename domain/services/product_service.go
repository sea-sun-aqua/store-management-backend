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

func (p *productService) UpdateProductByID(ctx context.Context, id string, req *requests.ProductUpdateAmountRequest) (*models.Product, error) {
	if req.Amount < 0 {
		return nil, exceptions.ErrInvalidAmount
	}

	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, exceptions.ErrProductNotFound
	}

	err = p.productRepo.UpdateAmountByID(ctx, id, req)
	if err != nil {
		return nil, err
	}

	product.Amount += req.Amount

	return product, nil
}

func (p *productService) FindByName(ctx context.Context, name string) (*models.Product, error) {
	return p.productRepo.FindByName(ctx, name)
}

func (p *productService) FindByID(ctx context.Context, id string) (*models.Product, error) {
	return p.productRepo.FindByID(ctx, id)
}

func (p *productService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	return p.productRepo.GetAll(ctx)
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
