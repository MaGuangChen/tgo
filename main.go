package main

import (
	"fmt"

	"github.com/G-Cool-ThanosGo/controller"
	"github.com/gin-gonic/gin"
)

var env = "develop"
var build = ""
var version = "1.0.0"

func main() {
	fmt.Printf("[%s] thanosGo version: %s [%s]!!!!!!\n", env, version, build)

	r := gin.Default()
	r.POST("/report/dodo", controller.DodoReport)
	r.Run(":5000")
}
