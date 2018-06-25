package iface

import (
	"github.com/G-Cool-ThanosGo/model/schema"
	"github.com/G-Cool-ThanosGo/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataBase interface {
	connect() *gorm.DB
}

type MySqlDB struct {
	Orders schema.Orders
}
type SqliteDB struct{}

func (MySqlDB) connect() *gorm.DB {
	db, err := gorm.Open("mysql", "fin_paul:3CzjWc#JY$i@hr@tcp(35.189.162.52:3306)/msqdbt1?charset=utf8&parseTime=true")
	util.CheckError(err)
	return db
}

func ConnectDB(d DataBase) *gorm.DB {
	db := d.connect()
	return db
}
