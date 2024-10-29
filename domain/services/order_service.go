package services

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/exceptions"
	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
)

type orderService struct {
	orderRepo repositories.OrderRepository
}


func NewOrderService(orderRepo repositories.OrderRepository) usecases.OrderUseCase {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) Create(ctx context.Context, req *requests.OrderCreateRequest) error {	
	result, err := o.orderRepo.FindByID(ctx, req.OrderID)
	if err != nil {
		return err
	}
	if result != nil {
		return exceptions.ErrDuplicatedIDOrder
	}

	return o.orderRepo.Create(ctx, req)
}

func (o *orderService) GetAll(ctx context.Context) ([]models.Order, error) {
	return o.orderRepo.GetAll(ctx)
}

func (o *orderService) UpdateStatusByID(ctx context.Context, id string, req *requests.OrderUpdateStatusRequest) (*models.Order, error) {
	if !checkOrderStatus(req.Status) {
		return nil, exceptions.ErrStatusInvalid
	}
	order, err := o.orderRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, exceptions.ErrOrderIDNotFound
	}

	err = o.orderRepo.UpdateStatusByID(ctx, id, req)
	if err != nil {
		return nil, err
	}

	// update order
	order.Status = req.Status
	return order, nil
}

func checkOrderStatus(status string) bool {
    switch status {
    case models.OrderStatusTranferring, models.OrderStatusPack, models.OrderStatusDeliver, models.OrderStatusCompleted:
        return true
    default:
        return false
    }
}
