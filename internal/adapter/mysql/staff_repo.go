package mysql

import (
	"context"
	"database/sql"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type staffMYSQLRepository struct {
	db *sqlx.DB
}

func NewStaffMYSQLRepository(db *sqlx.DB) repositories.StaffRepository {
	return &staffMYSQLRepository{
		db: db,
	}
}

// Create implements repositories.StaffRepository.
func (s *staffMYSQLRepository) Create(ctx context.Context, req *requests.StaffRegisterRequest) error {
	// Generate UUID
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	_, err = s.db.QueryContext(ctx, "INSERT INTO staff (staff_id, staff_name, staff_email, staff_password) VALUES (?, ?, ?, ?)", id.String(), req.Name, req.Email, req.Password)

	return err
}

// FindByEmail implements repositories.StaffRepository.
func (s *staffMYSQLRepository) FindByEmail(ctx context.Context, email string) (*models.Staff, error) {
		var staff models.Staff
		err := s.db.QueryRowContext(ctx, "SELECT staff_id, staff_name, staff_email, staff_password FROM staff WHERE staff_email = ?", email).Scan(&staff.StaffID, &staff.Name, &staff.Email, &staff.Password)
		if err == sql.ErrNoRows {
			return nil, nil
		}
	
		if err != nil {
			return nil, err
		}
	
		return &staff, nil
}
