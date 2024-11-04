package usecases

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/responses"
)

type OrderUseCase interface {
	Create(ctx context.Context, req *requests.OrderCreateRequest) error
	UpdateStatusByID(ctx context.Context, id string, req *requests.OrderUpdateStatusRequest) (*models.Order, error)
	GetAll(ctx context.Context) ([]responses.Order, error)
}