package main

import (
	"examples/database"
	"github.com/kataras/iris/v12"
	"github.com/rinatusmanov/gorestfull/drivers/iris_driver"
	"github.com/rinatusmanov/gorestfull/maker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	app := iris.New()
	dsn := os.Getenv("dsn")
	dbGlobal, errDbGlobal := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if errDbGlobal != nil {
		panic(errDbGlobal)
	}
	dbGlobal.
		AutoMigrate(&database.Log{})

	cr, _ := maker.Maker(iris_driver.NewDriver(app), dbGlobal)
	cr.Crud(database.Log{})
	panic(app.Listen(":3000"))
}
