package models

import (
	"fmt"

	"github.com/k1/go-blog/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID        uint32 `gorm:"primary_key" json:"id"`
	CreatedAt uint32 `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt uint32 `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
	DeletedAt uint32 `json:"deleted_at"`
}

const connectDSN = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func NewDBEngine(config *configs.Mysql) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(connectDSN, config.User, config.Password, config.Host, config.Port, config.DBName)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdelConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	return db, nil
}
