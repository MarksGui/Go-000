package model

import (
	"Go-000/Week04/configs"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBEngine(conf *configs.Config) (*gorm.DB, error) {
	db, err := gorm.Open(conf.Database.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		conf.Database.UserName,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.DBName,
		conf.Database.Charset,
		conf.Database.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}
