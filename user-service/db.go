package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/jinzhu/gorm"
	"blog-micro/user-service/config"
	"github.com/kataras/iris/core/errors"
)

var (
	DB        *gorm.DB
)

func CreateConnection(conf config.Config) (*gorm.DB, error) {
	dbUser := conf.DB.User
	dbPass := conf.DB.Password
	dbHost := conf.DB.Host
	dbName := conf.DB.Name
	dbPort := fmt.Sprintf("%d", conf.DB.Port)

	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost,
			dbPort, dbName))
	if err != nil {
		return nil, errors.New("connect database failed " + err.Error())
	}
	return db, nil
}