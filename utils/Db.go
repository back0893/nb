package utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDb() (map[string]*gorm.DB, error) {
	dbs := make(map[string]*gorm.DB)
	for key, dbConfig := range GlobalObject.Database {
		config := mysql.NewConfig()
		config.User = dbConfig.User
		config.Passwd = dbConfig.Passwd
		config.Addr = dbConfig.Addr
		config.Net = dbConfig.Net
		config.Params = dbConfig.Params
		config.DBName = dbConfig.DbName
		db, err := gorm.Open("mysql", config.FormatDSN())
		if err != nil {
			return nil, err
		}
		dbs[key] = db
	}
	return dbs, nil
}
