package apis

import (
	"database/sql"
	"fmt"

	"github.com/G-Cool-ThanosGo/iface"
	"github.com/G-Cool-ThanosGo/service"
	"github.com/G-Cool-ThanosGo/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type exportType struct {
	tPlus    string // T + 1 or T + 4 or MT + 4
	upload   bool
	sendMail bool
	date     string
}

// DoReport : 給嘟嘟房的報表所需
func DoReport(c *gin.Context, log util.LogStruct, mysqldb iface.MySQLDB, gdb *gorm.DB, rawDB *sql.DB) {
	start, end, filePath, purpose := service.HandleReportParms(c)
	log.Info("[Controller][Report]", purpose, " : ", "\nstart: ", start, "\nend: ", end, "\nfilePath: ", filePath, "\npurpose: ", purpose)

	pkgOrders, pkgExtraction := mysqldb.PkgOperator.GetByParkTime(rawDB, start, end)

	fmt.Println("This is orders count: ", len(pkgOrders))
	mysqldb.PkgOperator.GetPkgFinance(pkgExtraction, gdb)
	// fmt.Println("This is park info count: ", len(pkgExtraction.OrdersID))
}

// var parkingRecordDetails []schema.ParkingRecordDetails
// db.Find(&parkingRecordDetails, "order_id in (?)", []int{orders[0].ID})

// fmt.Println("this is CaptureTime: ", o[0].CaptureTime)
// fmt.Println("this is parkingRecordDetails: ", parkingRecordDetails[0].ParkingRecordID)
// parkingRecord := mysqldb.ParkingRecordOperator.GetByExitedAt(db, start, end)
// parkingRecordDetails := mysqldb.ParkingRecordDetailOperator.GetByParkingRecordID(db, []int{33})

// orders := mysqldb.OrderOperator.GetByID(db, 1679)
// fmt.Println("this is CaptureTime: ", orders[0].ID)
// parkingRecordOperator := dboperator.ParkingRecordOperator{}
// parkingRecord := parkingRecordOperator.GetByExitedAt(db, start, end)
// fmt.Println("this is parkingRecord: ", parkingRecord)
// fmt.Println("this is parkingRecord: ", parkingRecordDetails)

// // manulOrders := report.mysqldb.OrderOperator.GetT4ManulAndRefund(report.rawDB, report.start, report.end, report.ordersID)
// fmt.Println(len(report.orders))
// fmt.Println(len(manulOrders))

// PKGReport :
// type PKGReport interface {
// 	HandleReportParms(c *gin.Context)
// }

// // Dodo : 給嘟嘟房的報表
// type Dodo struct{}

// // Finance :
// type Finance struct{}
