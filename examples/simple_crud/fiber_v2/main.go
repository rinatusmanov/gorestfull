package main

import (
	"examples/database"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatusmanov/gorestfull/drivers/fiber_v2_driver"
	"github.com/rinatusmanov/gorestfull/maker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	app := fiber.New()
	driver := fiber_v2_driver.NewDriver(app)
	dsn := os.Getenv("dsn")
	dbGlobal, errDbGlobal := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if errDbGlobal != nil {
		panic(errDbGlobal)
	}
	dbGlobal.
		AutoMigrate(&database.Log{})
	result, errCrud := maker.Maker(driver, dbGlobal)
	if errCrud == nil {
		result.Crud(database.Log{})
	}
	panic(app.Listen(":3000"))
}
