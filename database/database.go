package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mojachieee/go-HoneyPot/config"
)

// InitDatabase is used to initialise the database connection and returns a pointer to the db
func InitDatabase(cfg config.Database) *gorm.DB {
	port := cfg.Port
	if port == "" {
		port = "3306"
	}
	str := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, port, cfg.Name)
	db, err := gorm.Open("mysql", str)
	if err != nil {
		log.Fatal(err)
	}
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
