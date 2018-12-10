package iface

import (
	"database/sql"

	"github.com/G-Cool-ThanosGo/app"
	"github.com/G-Cool-ThanosGo/models/dboperator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DataBase : db interface
type DataBase interface {
	connect() *gorm.DB
}

// MySQLDB : MySQLDB
type MySQLDB struct {
	BanksOperator              dboperator.BanksOperator
	CreditCardsOperator        dboperator.CreditCardsOperator
	GatewaysOperator           dboperator.GatewaysOperator
	InvoicesOperator           dboperator.InvoicesOperator
	InvitationCode             dboperator.InvitationCodeOperator
	MemberPointRedeemsOperator dboperator.MemberPointRedeemsOperator
	PkgOperator                dboperator.PkgOperator
	PaymentDetailsOperator     dboperator.PaymentDetailsOperator
}

// SqliteDB : SqliteDB
type SqliteDB struct{}

func (MySQLDB) connect() *gorm.DB {
	db, err := gorm.Open("mysql", "paul:5678@tcp(xx.xxx.xxx.xx:3306)/db1?charset=utf8&parseTime=true")
	app.CheckError(err)
	return db
}

// ConnectDB : 連接DB使用gorm
func ConnectDB(d DataBase) *gorm.DB {
	db := d.connect()
	return db
}

// ConnectDBUseRawSQL : 連接DB使用sql package
func ConnectDBUseRawSQL() *sql.DB {
	rawDB, connectErr := sql.Open("mysql", "paul:5678@tcp(xx.xxx.xxx.xx:3306)/db1?charset=utf8&parseTime=true")
	app.CheckError(connectErr)
	connectErr = rawDB.Ping()

	return rawDB
}
