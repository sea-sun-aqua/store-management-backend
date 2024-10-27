package services

import (
	"context"

	"github.com/FLUKKIES/marketplace-backend/domain/exceptions"
	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/usecases"
	"golang.org/x/crypto/bcrypt"
)

type staffService struct {
	staffRepo repositories.StaffRepository
}


func NewStaffService(staffRepo repositories.StaffRepository) usecases.StaffUseCase {
	return &staffService{
		staffRepo: staffRepo,
	}
}

func (s *staffService) Login(ctx context.Context, req *requests.StaffLoginRequest) (*models.Staff, error) {
	staff, err := s.staffRepo.FindByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	// Check if staff exist
	if staff == nil {	
		return nil, exceptions.ErrLoginFailed
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(req.Password)) != nil {
		return nil, exceptions.ErrLoginFailed
	}
	


	return staff, nil
}


func (s *staffService) Register(ctx context.Context, req *requests.StaffRegisterRequest) error {
	staff, err := s.staffRepo.FindByEmail(ctx, req.Email)

	if err != nil {
		return err
	}

	// Check if staff already exist
	if staff != nil {
		return exceptions.ErrDuplicatedEmail
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)
	
	return s.staffRepo.Create(ctx, req)
}
