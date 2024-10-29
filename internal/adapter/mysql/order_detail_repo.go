package mysql

import (
	"context"
	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/jmoiron/sqlx"
)

type orderDetailMYSQLRepository struct {
	db *sqlx.DB
}

func NewOrderDetailMYSQLRepository(db *sqlx.DB) repositories.OrderDetailRepository {
	return &orderDetailMYSQLRepository{
		db: db,
	}
}

func (o *orderDetailMYSQLRepository) FindByOrderID(ctx context.Context, id string) ([]models.OrderDetail, error) {
	var orderDetails []models.OrderDetail
	err := o.db.SelectContext(ctx, &orderDetails, `SELECT * FROM order_detail WHERE order_id = ?`, id)
	if err != nil {
		return nil, err
	}
	return orderDetails, nil
}

func (o *orderDetailMYSQLRepository) FindByProductID(ctx context.Context, id string) ([]models.OrderDetail, error) {
	var orderDetails []models.OrderDetail
	err := o.db.SelectContext(ctx, &orderDetails, `SELECT * FROM order_detail WHERE product_id = ?`, id)
	if err != nil {
		return nil, err
	}
	return orderDetails, nil
}


func (o *orderDetailMYSQLRepository) GetAll(ctx context.Context) ([]models.OrderDetail, error) {
	var orderDetails []models.OrderDetail
	err := o.db.SelectContext(ctx, &orderDetails, `SELECT * FROM order_detail`)
	if err != nil {
		return nil, err
	}
	return orderDetails, nil
}

