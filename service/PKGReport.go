package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// PKGRequestParms : 請求PKG報表時所帶入的參數
type PKGRequestParms struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	FilePath string `json:"filePath"`
	Purpose  string `json:"purpose"`
}

// HandleReportParms : 處理PKG參數
func HandleReportParms(c *gin.Context) (time.Time, time.Time, string, string) {
	r := PKGRequestParms{}
	if c.BindJSON(&r) != nil {
		fmt.Println("[handleReportParms] fail")
	}
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02 15:04:05"
	start, _ := time.ParseInLocation(timeLayout, r.Start, loc)
	end, _ := time.ParseInLocation(timeLayout, r.End, loc)
	filePath := r.FilePath
	purpose := r.Purpose
	return start, end, filePath, purpose
}
