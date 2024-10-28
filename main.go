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
	"github.com/jmoiron/sqlx"
)

func main() {
	app := fiber.New()


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

	app.Post("/register", staffHandler.Register) //correct
	app.Post("/login", staffHandler.Login) //correct

	app.Post("/product", productHandler.Register) //correct
	app.Post("/product/:ProductID", productHandler.UpdateProductByID) //correct
	app.Get("/product/:ProductID", productHandler.FindByID) //correct
	app.Get("/product", productHandler.GetAllProducts) //correct

	app.Post("/order", orderHandler.Create) //correct
	app.Post("/order/:OrderID", orderHandler.UpdateStatusOrder) //correct
	app.Get("/order", orderHandler.GetAll) //correct


	//ทำงานก่อน return 
	defer db.Close()

	// สามารถใส่ใน env ได้
	if err := app.Listen(":9000"); err != nil {
		log.Fatal(err)
	}
}