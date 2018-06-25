package controller

import (
	"fmt"

	"github.com/G-Cool-ThanosGo/iface"
	"github.com/G-Cool-ThanosGo/model/schema"
	"github.com/G-Cool-ThanosGo/util"
	"github.com/gin-gonic/gin"
)

type exportType struct {
	tPlus    string // T + 1 or T + 4 or MT + 4
	upload   bool
	sendMail bool
	date     string
}

type PKGRequestParms struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	FilePath string `json:"filePath"`
	Purpose  string `json:"purpose"`
}

type dodo struct{}
type fiance struct{}

func DodoReport(c *gin.Context) {
	// step 1 處理參數
	req := PKGRequestParms{}
	start, end, filePath, purpose := handleReportParms(c, req)
	init := util.InitUtil{}
	log := init.LogInit()
	log.Info("[Controller][Report][Dodo]: ", "\nstart: ", start, "\nend: ", end, "\nfilePath: ", filePath, "\npurpose: ", purpose)

	// step 2 連線至mysql
	mysqldb := iface.MySqlDB{}
	db := iface.ConnectDB(mysqldb)

	var orders []schema.Orders
	db.Where("id = ?", 375).Find(&orders)
	fmt.Println("this is CreatedAt: ", orders[0].CreatedAt)
}

func handleReportParms(c *gin.Context, r PKGRequestParms) (string, string, string, string) {
	if c.BindJSON(&r) != nil {
		fmt.Println("[handleReportParms] fail")
	}
	start := r.Start
	end := r.End
	filePath := r.FilePath
	purpose := r.Purpose
	return start, end, filePath, purpose
}
