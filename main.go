package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FLUKKIES/marketplace-backend/configs"
	"github.com/FLUKKIES/marketplace-backend/domain/services"
	"github.com/FLUKKIES/marketplace-backend/internal/adapter/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/FLUKKIES/marketplace-backend/internal/adapter/rest"
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
	staffRepo := mysql.NewStaffMYSQLRepository(db)
	staffService := services.NewStaffService(staffRepo)
	staffHandler := rest.NewStaffHandler(staffService)

	app.Post("/register", staffHandler.Register)
	app.Post("/login", staffHandler.Login)

	//ทำงานก่อน return 
	defer db.Close()

	// สามารถใส่ใน env ได้
	if err := app.Listen(":9000"); err != nil {
		log.Fatal(err)
	}
}