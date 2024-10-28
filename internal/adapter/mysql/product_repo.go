package mysql

import (
	"context"
	"database/sql"

	"github.com/FLUKKIES/marketplace-backend/domain/models"
	"github.com/FLUKKIES/marketplace-backend/domain/repositories"
	"github.com/FLUKKIES/marketplace-backend/domain/requests"
	"github.com/jmoiron/sqlx"
)

type productMYSQLRepository struct {
	db *sqlx.DB
}


func NewProductMYSQLRepository(db *sqlx.DB) repositories.ProductRepository {
	return &productMYSQLRepository{
		db: db,
	}
}

func (p *productMYSQLRepository) UpdateAmountByID(ctx context.Context, id string, req *requests.ProductUpdateAmountRequest) error {
	_, err := p.db.ExecContext(ctx, "UPDATE product SET product_amount = product_amount + ? WHERE product_id = ?", req.Amount, id)
	return err
}

func (p *productMYSQLRepository) Create(ctx context.Context, req *requests.ProductRegisterRequest) error {
	_, err := p.db.QueryContext(ctx, "INSERT INTO product (product_id, product_name, product_price, product_amount, safety_stock_amount) VALUES (?, ?, ?, ?, ?)", req.ProductID, req.Name, req.PriceOfUnit, req.Amount, req.QuantityOfSafetyStock)
	return err
}

func (p *productMYSQLRepository) FindByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	err := p.db.QueryRowContext(ctx, "SELECT product_id, product_name, product_price, product_amount, safety_stock_amount FROM product WHERE product_id = ?", id).Scan(&product.ProductID, &product.Name, &product.Price, &product.Amount, &product.QuantityOfSafetyStock)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	println(product.ProductID, product.Name, product.Price, product.Amount, product.QuantityOfSafetyStock)

	return &product, nil
}

func (p *productMYSQLRepository) FindByName(ctx context.Context, name string) (*models.Product, error) {
	var product models.Product
	err := p.db.QueryRowContext(ctx, "SELECT product_id, product_name, product_price, product_amount, safety_stock_amount FROM product WHERE product_name = ?", name).Scan(&product.ProductID, &product.Name, &product.Price, &product.Amount, &product.QuantityOfSafetyStock)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productMYSQLRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	rows, err := p.db.Query("SELECT * FROM product")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Amount, &p.QuantityOfSafetyStock)
		if err != nil {
			panic(err.Error())

		}
		products = append(products, p)
	}

	return products, nil
}
