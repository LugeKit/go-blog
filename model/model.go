package model

import (
	"fmt"
	"time"

	"github.com/k1/go-blog/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

const (
	dbConnectString = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local"
	maxIdleConns    = 10
	maxOpenConns    = 100
	connMaxLifetime = time.Hour
)

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

func init() {
	var (
		err                                             error
		dbName, user, password, host, port, tablePrefix string
	)

	dbName = conf.Conf.Database.Mysql.DBName
	user = conf.Conf.Database.Mysql.User
	password = conf.Conf.Database.Mysql.Password
	host = conf.Conf.Database.Mysql.Host
	port = conf.Conf.Database.Mysql.Port
	tablePrefix = conf.Conf.Database.Mysql.TablePrefix

	db, err = gorm.Open(mysql.Open(fmt.Sprintf(
		dbConnectString,
		user,
		password,
		host,
		port,
		dbName,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	setConnParam()
}

func setConnParam() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)
}
