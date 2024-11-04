package mysql

import (
	"context"
	"database/sql"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/FLUKKIES/marketplace-backend/domain/responses"
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

func (o *orderMYSQLRepository) GetAll(ctx context.Context) ([]responses.Order, error) {
	var response []responses.Order

	var orders []models.Order
	err := o.db.SelectContext(ctx, &orders, "SELECT order_id, customer_name, customer_email, phone_number, customer_address, order_payment_method, order_status, order_created_date, staff_id FROM orders")
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		var staff responses.Staff
		err = o.db.GetContext(ctx, &staff, "SELECT staff_id, staff_name, staff_email FROM staff WHERE staff_id = ?", order.StaffID)
		if err != nil {
			return nil, err
		}

		var products []responses.Product

		err = o.db.SelectContext(ctx, &products, `
			SELECT 
    			p.product_id AS product_id,
    			p.product_name AS product_name,
    			p.product_price AS product_price,
    			o.order_amount AS order_amount
			FROM 
    			product p
			JOIN 
    			order_detail o ON p.product_id = o.product_id
			WHERE 
    			o.order_id = ?;
		`, order.OrderID)
		
		if err != nil {
			return nil, err
		}


		// สร้าง object แล้วใส่ list
		res := responses.Order{
			OrderID: order.OrderID,
			CustomerName: order.CustomerName,
			Email: order.Email,
			PhoneNumber: order.PhoneNumber,
			Address: order.Address,
			Method:order.Method,
			Status: order.Status,
			CreatedAt: order.CreatedAt,
			Staff: staff,
			Products: products,
		}
		response = append(response, res)

	}
	
	return response, nil
}

func (o *orderMYSQLRepository) UpdateStatusByID(ctx context.Context, id string, req *requests.OrderUpdateStatusRequest) error {
	_, err := o.db.ExecContext(ctx, "UPDATE orders SET order_status = ? WHERE order_id = ?", req.Status, id)
	if(req.Status == models.OrderStatusCompleted){
		for _, product := range req.Products {
			_, err := o.db.ExecContext(ctx, "UPDATE product SET product_amount = product_amount - ? WHERE product_id = ?", product.Amount, product.ProductID)
			if err != nil {
				return err
			}
		}
	}
	return err
}

