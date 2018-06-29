package main

import (
	"fmt"

	"github.com/G-Cool-ThanosGo/apis"
	"github.com/G-Cool-ThanosGo/iface"
	"github.com/G-Cool-ThanosGo/util"
	"github.com/gin-gonic/gin"
)

var env = "develop"
var build = ""
var version = "1.0.0"

// Controllers : 流程控制
// type Controllers struct {
// 	log        *util.LogStruct
// 	mysqldb    iface.MySQLDB
// 	db         *gorm.DB
// 	rawDB      *sql.DB
// 	dodoReport app.Report
// }

func main() {
	fmt.Println("Start Server ThanosGo version:", env, version, build)
	r := gin.Default()
	// log := util.LogInit()
	log := util.LogStruct{}
	log.Config("127.0.0.1:5000")
	mysqldb := iface.MySQLDB{}
	gdb := iface.ConnectDB(mysqldb)
	rawDB := iface.ConnectDBUseRawSQL()

	r.POST("/report/dodo", func(c *gin.Context) { apis.DoReport(c, log, mysqldb, gdb, rawDB) })
	r.Run(":5000")
}
