package controller

import (
	"fmt"

	"github.com/G-Cool-ThanosGo/iface"
	"github.com/G-Cool-ThanosGo/util"
	"github.com/gin-gonic/gin"
)

type exportType struct {
	tPlus    string // T + 1 or T + 4 or MT + 4
	upload   bool
	sendMail bool
	date     string
}

type dodoRequestParms struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	FilePath string `json:"filePath"`
	Purpose  string `json:"purpose"`
}

type dodo struct{}
type fiance struct{}

func DodoReport(c *gin.Context) {
	req := dodoRequestParms{}
	if c.BindJSON(&req) != nil {
		return
	}
	start, end, filePath, purpose := handleReportParms(req)

	init := util.InitUtil{}
	log := init.LogInit()
	log.Info("[Controller][Report][Dodo]: ", req)
	// step 1 處理參數
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(filePath)
	fmt.Println(purpose)

	// step 2 根據參數判斷連線database
	msqdbt1 := iface.MySqlDB{}
	iface.ConnectDB(msqdbt1)
}

func handleReportParms(r dodoRequestParms) (string, string, string, string) {
	start := r.Start
	end := r.End
	filePath := r.FilePath
	purpose := r.Purpose
	return start, end, filePath, purpose
}

// type report interface {
// 	handleParms()
// 	getOrder()
// }
// type report struct {
// 	log util.LogStruct
// 	db  iface.DataBase
// }
