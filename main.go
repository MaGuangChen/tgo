package main

import (
	"fmt"

	"github.com/G-Cool-ThanosGo/apis"
	"github.com/G-Cool-ThanosGo/app"
	"github.com/G-Cool-ThanosGo/iface"
	"github.com/gin-gonic/gin"
)

var env = "develop"
var build = ""
var version = "1.0.0"

func main() {
	fmt.Println("Start Server ThanosGo version:", env, version, build)
	r := gin.Default()
	log := app.LogStruct{}
	log.Config("127.0.0.1:5000")
	mysqldb := iface.MySQLDB{}
	gdb := iface.ConnectDB(mysqldb)
	rawDB := iface.ConnectDBUseRawSQL()

	r.POST("/report/dodo", func(c *gin.Context) { apis.DoReport(c, log, mysqldb, gdb, rawDB) })
	r.Run(":5000")
}

// Controllers : 流程控制
// type Controllers struct {
// 	log        *util.LogStruct
// 	mysqldb    iface.MySQLDB
// 	db         *gorm.DB
// 	rawDB      *sql.DB
// 	dodoReport app.Report
// }

// links := []string{
// 	"http://facebook.com",
// 	"http://stackoverflow.com",
// 	"http://golang.org",
// 	"http://amazon.com",
// 	"http://google.com",
// }
// l0 := "http://facebook.com"
// l1 := "http://google.com"
// go checkL(l0)
// go checkL(l1)

// channel is type and use to comunicate with other routines
// c := make(chan string)

// for _, link := range links {
// 	go checkLink(link, c)
// }

// func checkL(l string) {
// 	_, err := http.Get(l) // blocking code
// 	if err != nil {
// 		fmt.Println(l, "might be down")
// 		return
// 	}

// 	fmt.Println(l, "is up")
// }

// func checkLink(link string, c chan string) {
// 	// time.Sleep(5 * time.Second) // pause the 5 second
// 	_, err := http.Get(link) // blocking code
// 	// we put return here because we want to make sure we dont do any things
// 	// in err case
// 	if err != nil {
// 		fmt.Println(link, "might be down")
// 		c <- link
// 		return
// 	}

// 	fmt.Println(link, "is up")
// 	c <- link
// }
