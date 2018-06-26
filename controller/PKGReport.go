package controller

import (
	"github.com/G-Cool-ThanosGo/iface"
	"github.com/G-Cool-ThanosGo/service"
	"github.com/G-Cool-ThanosGo/util"
	"github.com/gin-gonic/gin"
)

type exportType struct {
	tPlus    string // T + 1 or T + 4 or MT + 4
	upload   bool
	sendMail bool
	date     string
}

type dodo struct{}
type fiance struct{}

func DodoReport(c *gin.Context) {
	// step 1 處理參數
	start, end, filePath, purpose := service.HandleReportParms(c)
	init := util.InitUtil{}
	log := init.LogInit()
	log.Info("[Controller][Report][Dodo]: ", "\nstart: ", start, "\nend: ", end, "\nfilePath: ", filePath, "\npurpose: ", purpose)

	// step 2 連線至mysql
	mysqldb := iface.MySqlDB{}
	// db := iface.ConnectDB(mysqldb)
	// defer db.Close()
	rawDB := iface.ConnectDBUseRawSql()
	defer rawDB.Close()
	// parkingRecord := mysqldb.ParkingRecordOperator.GetByExitedAt(db, start, end)
	// parkingRecordDetails := mysqldb.ParkingRecordDetailOperator.GetByParkingRecordID(db, []int{33})

	// orders := mysqldb.OrderOperator.GetByID(db, 1679)
	// fmt.Println("this is CaptureTime: ", orders[0].ID)

	mysqldb.OrderOperator.GetByParkTime(rawDB, start, end)
	// fmt.Println("this is CaptureTime: ", orders)

	// parkingRecordOperator := dboperator.ParkingRecordOperator{}
	// parkingRecord := parkingRecordOperator.GetByExitedAt(db, start, end)
	// fmt.Println("this is parkingRecord: ", parkingRecord)
	// fmt.Println("this is parkingRecord: ", parkingRecordDetails)
}

// var parkingRecordDetails []schema.ParkingRecordDetails
// db.Find(&parkingRecordDetails, "order_id in (?)", []int{orders[0].ID})

// fmt.Println("this is CaptureTime: ", o[0].CaptureTime)
// fmt.Println("this is parkingRecordDetails: ", parkingRecordDetails[0].ParkingRecordID)
