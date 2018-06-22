package iface

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataBase interface {
	connect()
}

type MySqlDB struct{}
type Sqlite3DB struct {
	file string
}

func (MySqlDB) connect() {
	db, err := gorm.Open("mysql", "fin_paul:3CzjWc#JY$i@hr@tcp(35.189.162.52:3306)/msqdbt1?charset=utf8")
	if err != nil {
		fmt.Println("err is: ", err)
		return
	}
	defer db.Close()
}

func (Sqlite3DB) connect(file Sqlite3DB) {
	db, err := gorm.Open("sqlite3", file)
	if err != nil {
		fmt.Println("err is: ", err)
		return
	}
	defer db.Close()
}

func ConnectDB(d DataBase) {
	d.connect()
}
