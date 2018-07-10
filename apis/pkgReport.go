package apis

import (
	"database/sql"
	"fmt"

	"github.com/G-Cool-ThanosGo/app"
	"github.com/G-Cool-ThanosGo/iface"
	"github.com/G-Cool-ThanosGo/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type exportType struct {
	tPlus    string // T + 1 or T + 4 or MT + 4
	upload   bool
	sendMail bool
	date     string
}

// DoReport : 嘟嘟房、會計 報表 park_detail, pref_detail
func DoReport(c *gin.Context, log app.LogStruct, mysqldb iface.MySQLDB, gdb *gorm.DB, rawDB *sql.DB) {
	// 處理請求參數
	start, end, filePath, purpose := service.HandleReportParms(c)
	log.Info("[Controller][Report]", purpose, " : ", "\nstart: ", start, "\nend: ", end, "\nfilePath: ", filePath, "\npurpose: ", purpose)

	// 取得 pkg orders OrdersID AccountID LotsID ParkingRecordDetailsID
	pkgOrders, pkgExtraction := mysqldb.PkgOperator.GetByParkTime(rawDB, start, end)
	fmt.Println("This is orders in DoReport count: ", len(pkgOrders))

	// 透過 order prop 取得 paymentDetails, invoices, invitationCode, memberPointRedeems, orderCreateRecords, orderModRecords
	pkgDataByOrdersID := mysqldb.PkgOperator.GetByOrdersID(pkgExtraction, gdb)
	fmt.Println("this is payment details length: ", len(pkgDataByOrdersID.PaymentDetails))

	// 透過 parking_record_details prop 寫 raw sql
	mysqldb.PkgOperator.GetByParkingRecordDetailsID(pkgExtraction, gdb)
	mysqldb.PkgOperator.GetByPaymentDetailsID(pkgDataByOrdersID.PaymentDetailsID, gdb)
	// 透過payment_details prop 寫 raw sql
}

// oidChan := make(chan dboperator.GetPkgDataByOrdersID)
// go func() {
// pkgDataByOrdersID := mysqldb.PkgOperator.GetByOrdersID(pkgExtraction, gdb)
// 	oidChan <- pkgDataByOrdersID
// }()
// pkgDataByOrdersID := <-oidChan
