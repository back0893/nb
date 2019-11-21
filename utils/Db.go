package utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDb() (*gorm.DB, error) {
	config := mysql.NewConfig()
	config.User = GlobalObject.Database.User
	config.Passwd = GlobalObject.Database.Passwd
	config.Addr = GlobalObject.Database.Addr
	config.Net = GlobalObject.Database.Net
	config.Params = GlobalObject.Database.Params
	config.DBName = GlobalObject.Database.DbName
	return gorm.Open("mysql", config.FormatDSN())
}
