package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IMysql interface {
	AutoMigrate(model interface{}) error
	GetInstance() *gorm.DB
}

type conn struct {
	db *gorm.DB
}

func NewConnectMysql(user string, pass string, dbName string, host string, port string) (IMysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &conn{
		db: db,
	}, err
}

func (c conn) GetInstance() *gorm.DB {
	return c.db
}

func (c conn) AutoMigrate(model interface{}) error {
	c.db.AutoMigrate(
		&model,
	)
	return nil
}
