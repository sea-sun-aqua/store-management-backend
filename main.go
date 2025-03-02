package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FLUKKIES/marketplace-backend/configs"
	"github.com/FLUKKIES/marketplace-backend/domain/services"
	"github.com/FLUKKIES/marketplace-backend/internal/adapter/mysql"
	"github.com/FLUKKIES/marketplace-backend/internal/adapter/rest"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // อนุญาตให้ request มาจากโดเมนนี้
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))


	ctx := context.Background()

	cfg := configs.ReadConfig()

	db, err := sqlx.ConnectContext(ctx, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	if(err != nil){
		log.Fatal(err)
	}

	// Dependencies Injection
	// staff
	staffRepo := mysql.NewStaffMYSQLRepository(db)
	staffService := services.NewStaffService(staffRepo)
	staffHandler := rest.NewStaffHandler(staffService)

	// product
	productRepo := mysql.NewProductMYSQLRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := rest.NewProductHandler(productService)

	//orders
	orderRepo := mysql.NewOrderMYSQLRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := rest.NewOrderHandler(orderService)

	//order_detail
	orderDetailRepo := mysql.NewOrderDetailMYSQLRepository(db)
	orderDetailService := services.NewOrderDetailService(orderDetailRepo)
	orderDetailHandler := rest.NewOrderDetailHandler(orderDetailService)

	//purchase_order
	purchaseOrderRepo := mysql.NewPurchaseOrderMYSQLRepository(db)
	purchaseOrderService := services.NewPurchaseOrderService(purchaseOrderRepo)
	purchaseOrderHandler := rest.NewPurchaseOrderHandler(purchaseOrderService)

	//purchase-order-detail
	purchaseOrderDetailRepo := mysql.NewPurchaseOrderDetailMYSQLRepository(db)
	purchaseOrderDetailService := services.NewPurchaseOrderDetailService(purchaseOrderDetailRepo)
	purchaseOrderDetailHandler := rest.NewPurchaseOrderDetailHandler(purchaseOrderDetailService)


	app.Post("/register", staffHandler.Register) //correct
	app.Post("/login", staffHandler.Login) //correct

	app.Post("/product", productHandler.Register) //correct
	app.Post("/product/:ProductID", productHandler.UpdateProductByID) //correct
	app.Get("/product/:ProductID", productHandler.FindByID) //correct
	app.Get("/product", productHandler.GetAllProducts) //correct

	app.Post("/order", orderHandler.Create) //correct
	app.Post("/order/:OrderID", orderHandler.UpdateStatusOrder) //correct
	app.Get("/order", orderHandler.GetAll) //correct

	app.Get("/order-detail", orderDetailHandler.GetAll) //correct

	app.Post("/purchase-order", purchaseOrderHandler.Create) //correct
	app.Post("/purchase-order/:PurchaseOrderID", purchaseOrderHandler.UpdateStatusOrder) //correct
	app.Get("/purchase-order", purchaseOrderHandler.GetAll) //correct

	app.Get("/purchase-order-detail", purchaseOrderDetailHandler.GetAll)//correct


	//ทำงานก่อน return 
	defer db.Close()

	// สามารถใส่ใน env ได้
	if err := app.Listen(":9000"); err != nil {
		log.Fatal(err)
	}
}