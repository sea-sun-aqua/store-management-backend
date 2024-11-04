package usecases

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
)

type StaffUseCase interface {
	Register(ctx context.Context, req *requests.StaffRegisterRequest) error
	Login(ctx context.Context, req *requests.StaffLoginRequest) (*models.Staff, error)
}

// FindByID(ctx context.Context, id string) (*models.Staff, error)
// GetAll(ctx context.Context) ([]models.Staff, error)