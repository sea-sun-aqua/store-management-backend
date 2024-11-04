package mysql

import (
	"context"
	"database/sql"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/jmoiron/sqlx"
)

type purchaseOrderMYSQLRepository struct {
	db *sqlx.DB
}

func NewPurchaseOrderMYSQLRepository(db *sqlx.DB) repositories.PurchaseOrderRepository {
	return &purchaseOrderMYSQLRepository{
		db: db,
	}
}


func (p *purchaseOrderMYSQLRepository) Create(ctx context.Context, req *requests.PurchaseOrderCreateRequest) error {
	//insert into purchase_order table
	_, err := p.db.QueryContext(ctx,
		`INSERT INTO purchase_order (
		purchase_order_id, supplier_name, supplier_phone_number, purchase_order_status, purchase_order_created_date, purchase_order_payment_method, staff_id)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		req.PurchaseOrderID, req.SupplierName, req.PhoneNumber, req.Status, req.CreatedDate, req.Method, req.StaffID)
	if err != nil {
		return err
	}
	// insert into purchase_order_detail table
	products_req := req.Products
	for _, product := range products_req {
		_, err = p.db.QueryContext(ctx,
			`INSERT INTO purchase_order_detail (
			purchase_order_id, product_id, purchase_order_amount) VALUES (?, ?, ?)`,
			req.PurchaseOrderID, product.ProductID, product.Amount)
		if err != nil {
			return err
		}
	}
	return err
}

func (p *purchaseOrderMYSQLRepository) FindByID(ctx context.Context, id string) (*models.PurchaseOrder, error) {
	var purchaseOrder models.PurchaseOrder
	err := p.db.GetContext(ctx, &purchaseOrder, `SELECT * FROM purchase_order WHERE purchase_order_id = ?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &purchaseOrder, nil
}

func (p *purchaseOrderMYSQLRepository) GetAll(ctx context.Context) ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := p.db.SelectContext(ctx, &purchaseOrders, "SELECT * FROM purchase_order")
	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (p *purchaseOrderMYSQLRepository) UpdateStatusByID(ctx context.Context, id string, req *requests.PurchaseOrderUpdateStatusRequest) error {
	_, err := p.db.ExecContext(ctx, "UPDATE purchase_order SET purchase_order_status = ? WHERE purchase_order_id = ?", req.Status, id)
	return err
}
