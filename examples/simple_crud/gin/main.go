package main

import (
	"examples/database"
	"github.com/rinatusmanov/gorestfull/drivers/gin_driver"
	"github.com/rinatusmanov/gorestfull/maker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	dsn := os.Getenv("dsn")
	dbGlobal, errDbGlobal := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if errDbGlobal != nil {
		panic(errDbGlobal)
	}
	dbGlobal.
		AutoMigrate(&database.Log{})

	cr, _ := maker.Maker(gin_driver.NewDriver(r), dbGlobal)
	cr.Crud(database.Log{})
	panic(r.Run(":3000"))
}
