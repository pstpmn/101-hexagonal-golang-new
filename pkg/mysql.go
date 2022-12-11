package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql interface {
	AutoMigrate(model interface{}) error
}

type conn struct {
	db *gorm.DB
}

func NewConnect(user string, pass string, dbName string, host string, port string) (Mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &conn{
		db: db,
	}, err
}

func (c conn) AutoMigrate(model interface{}) error {
	c.db.AutoMigrate(
		&model,
	)
	return nil
}
