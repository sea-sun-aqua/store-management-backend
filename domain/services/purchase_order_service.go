package services

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/exceptions"
	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
)

type purchaseOrderService struct {
	purchaseOrderRepo repositories.PurchaseOrderRepository
}


func NewPurchaseOrderService(purchaseOrderRepo repositories.PurchaseOrderRepository) usecases.PurchaseOrderUseCase {
	return &purchaseOrderService{
		purchaseOrderRepo: purchaseOrderRepo,
	}
}

func (p *purchaseOrderService) Create(ctx context.Context, req *requests.PurchaseOrderCreateRequest) error {
	result, err := p.purchaseOrderRepo.FindByID(ctx, req.PurchaseOrderID)
	if err != nil {
		return err
	}
	if result != nil {
		return exceptions.ErrDuplicatedPurchaseIDOrder
	}

	return p.purchaseOrderRepo.Create(ctx, req)
}

func (p *purchaseOrderService) GetAll(ctx context.Context) ([]models.PurchaseOrder, error) {
	return p.purchaseOrderRepo.GetAll(ctx)
}

func (p *purchaseOrderService) UpdateStatusByID(ctx context.Context, id string, req *requests.PurchaseOrderUpdateStatusRequest) (*models.PurchaseOrder, error) {
	if !checkPurchaseOrderStatus(req.Status) {
		return nil, exceptions.ErrStatusInvalid
	}
	orderDetail, err := p.purchaseOrderRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if orderDetail == nil {
		return nil, exceptions.ErrPurchaseOrderIDNotFound
	}

	err = p.purchaseOrderRepo.UpdateStatusByID(ctx, id, req)
	if err != nil {
		return nil, err
	}

	// update orderDetail
	orderDetail.Status = req.Status
	return orderDetail, nil
}

func checkPurchaseOrderStatus(status string) bool {
	switch status {
	case models.PurchaseOrderStatusTranferring, models.PurchaseOrderStatusCanceled, models.PurchaseOrderStatusCompleted:
		return true
	default:
		return false
	}
}
