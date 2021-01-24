package main

import (
	"examples/database"
	"github.com/rinatusmanov/crud"
	"github.com/rinatusmanov/crud/drivers/http_driver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main() {
	// init *gorm.DB
	dsn := os.Getenv("dsn")
	dbGlobal, errDbGlobal := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if errDbGlobal != nil {
		panic(errDbGlobal)
	}
	dbGlobal.
		AutoMigrate(&database.Log{})

	cr, _ := crud.Maker(http_driver.NewDriver(http.DefaultServeMux), dbGlobal)
	cr.Crud(database.Log{})
	panic(http.ListenAndServe(":3000", nil))
}
