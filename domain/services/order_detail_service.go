package services

import (
	"context"
	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
)

type orderDetailService struct {
	orderDetailRepo repositories.OrderDetailRepository
}

func NewOrderDetailService(orderDetailRepo repositories.OrderDetailRepository) usecases.OrderDetailUseCase {
	return &orderDetailService{
		orderDetailRepo: orderDetailRepo,
	}
}

func (o *orderDetailService) FindByOrderID(ctx context.Context, id string) ([]models.OrderDetail, error) {
	return o.orderDetailRepo.FindByOrderID(ctx, id)
}

func (o *orderDetailService) FindByProductID(ctx context.Context, id string) ([]models.OrderDetail, error) {
	return o.orderDetailRepo.FindByProductID(ctx, id)
}

func (o *orderDetailService) GetAll(ctx context.Context) ([]models.OrderDetail, error) {
	return o.orderDetailRepo.GetAll(ctx)
}
