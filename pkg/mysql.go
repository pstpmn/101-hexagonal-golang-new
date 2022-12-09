package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql interface {
	Connect() (*gorm.DB, error)
}

type config struct {
	User   string
	Pass   string
	DbName string
	Host   string
	Port   string
}

func NewConnect(user string, pass string, dbName string, host string, port string) Mysql {
	return &config{User: user, Pass: pass, DbName: dbName, Host: host, Port: port}
}

func (c config) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Pass, c.Host, c.Port, c.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func AutoMigrate(conn *gorm.DB, model interface{}) error {
	conn.AutoMigrate(
		&model,
	)
	return nil
}
