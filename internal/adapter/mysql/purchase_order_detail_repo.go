package mysql

import (
	"context"
	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/jmoiron/sqlx"
)

type purchaseOrderDetailMYSQLRepository struct {
	db *sqlx.DB
}

func NewPurchaseOrderDetailMYSQLRepository(db *sqlx.DB) repositories.PurchaseOrderDetailRepository {
	return &purchaseOrderDetailMYSQLRepository{db: db}
}

func (p *purchaseOrderDetailMYSQLRepository) GetAll(ctx context.Context) ([]models.PurchaseOrderDetail, error) {
	var purchaseOrderDetails []models.PurchaseOrderDetail
	err := p.db.SelectContext(ctx, &purchaseOrderDetails, `SELECT * FROM purchase_order_detail`)
	if err != nil {
		return nil, err
	}
	return purchaseOrderDetails, nil
}
