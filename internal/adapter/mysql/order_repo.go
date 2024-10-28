package mysql

import (
	"context"
	"database/sql"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/jmoiron/sqlx"
)

type orderMYSQLRepository struct {
	db *sqlx.DB
}

func NewOrderMYSQLRepository(db *sqlx.DB) repositories.OrderRepository {
	return &orderMYSQLRepository{
		db: db,
	}
}


func (o *orderMYSQLRepository) Create(ctx context.Context, req *requests.OrderCreateRequest) error {
	// insert into orders table
	_, err := o.db.QueryContext(ctx, 
		`INSERT INTO orders (
		order_id, customer_name, customer_email, phone_number, customer_address, order_payment_method, order_status, order_created_date, staff_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, 
		req.OrderID, req.CustomerName, req.Email, req.PhoneNumber, req.Address, req.Method, req.Status, req.CreatedAt, req.StaffID)
	if err != nil {
		return err
	}


	// insert into order_detail table
	products_req := req.Products
	for _, p := range products_req {
		_, err = o.db.QueryContext(ctx,
			`INSERT INTO order_detail (
			order_id, product_id, order_amount) VALUES (?, ?, ?)`,
			req.OrderID, p.ProductID, p.Amount)
		if err != nil {
			return err
		}
	}
	return err
}

func (o *orderMYSQLRepository) FindByID(ctx context.Context, id string) (*models.Order, error) {
	var order models.Order
	//order_id, customer_name, customer_email, phone_number, customer_address, order_payment_method, order_status, order_created_date, staff_id
	err := o.db.GetContext(ctx, &order, `SELECT * FROM orders WHERE order_id = ?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *orderMYSQLRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	//order_id, customer_name, customer_email, phone_number, customer_address, order_payment_method, order_status, order_created_date, staff_id
	err := o.db.SelectContext(ctx, &orders, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderMYSQLRepository) UpdateStatusByID(ctx context.Context, id string, req *requests.OrderUpdateStatusRequest) error {
	_, err := o.db.ExecContext(ctx, "UPDATE orders SET order_status = ? WHERE order_id = ?", req.Status, id)
	return err
}

