package services

import (
	"context"
	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
)

type purchaseOrderDetailService struct {
	purchaseOrderDetailRepo repositories.PurchaseOrderDetailRepository
}

func NewPurchaseOrderDetailService(purchaseOrderDetailRepo repositories.PurchaseOrderDetailRepository) usecases.PurchaseOrderDetailUseCase {
	return &purchaseOrderDetailService{
		purchaseOrderDetailRepo: purchaseOrderDetailRepo,
	}
}

func (p *purchaseOrderDetailService) GetAll(ctx context.Context) ([]models.PurchaseOrderDetail, error) {
	return p.purchaseOrderDetailRepo.GetAll(ctx)
}
