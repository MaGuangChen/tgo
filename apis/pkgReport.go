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
	pkgService := service.PkgService{}
	start, end, filePath, purpose := pkgService.HandleReportParms(c)
	log.Info("[Controller][Report]", purpose, " : ", "\nstart: ", start, "\nend: ", end, "\nfilePath: ", filePath, "\npurpose: ", purpose)

	// 依照 PKG(停車出場時間) 向 DB 要求相關 data
	pkg := pkgService.GetPkgAllData(start, end, mysqldb, rawDB, gdb)
	fmt.Println(len(pkg.Orders))
	// 組合 PKG data
}

// oidChan := make(chan dboperator.GetPkgDataByOrdersID)
// go func() {
// pkgDataByOrdersID := mysqldb.PkgOperator.GetByOrdersID(pkgExtraction, gdb)
// 	oidChan <- pkgDataByOrdersID
// }()
// pkgDataByOrdersID := <-oidChan
