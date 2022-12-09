package blogApp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Application struct {
	Address string
}

var DB *gorm.DB

func (app *Application) Run() {
	r := gin.Default()
	BlogRouter(r)
	var err error
	DB, err = openDB("root", "abcdefg123!", "127.0.0.1", "3306", "blog")
	DB.AutoMigrate(&Blog{})
	if err != nil {
		log.Fatal(err)
	}
	err = r.Run(fmt.Sprintf(":%s", app.Address))
	if err != nil {
		log.Fatal(err)
	}
}
func openDB(name string, password string, url string, port string, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", name, password, url, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err

}
