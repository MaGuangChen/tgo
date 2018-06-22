package iface

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataBase interface {
	Connect()
}

type MySqlDB struct{}
type Sqlite3DB struct {
	file string
}

func (MySqlDB) Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "fin_paul:3CzjWc#JY$i@hr@tcp(35.189.162.52:3306)/msqdbt1?charset=utf8")
	if err != nil {
		fmt.Println("err is: ", err)
	} else {
		fmt.Println("connect msqdbt1 sucessed")
	}
	defer db.Close()

	return db
}

func (MySqlDB) Operate() {

}
