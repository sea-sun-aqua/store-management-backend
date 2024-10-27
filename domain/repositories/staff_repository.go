package repositories

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
)

type StaffRepository interface {
	Create(ctx context.Context, req *requests.StaffRegisterRequest) error
	FindByEmail(ctx context.Context, email string) (*models.Staff, error)
}
