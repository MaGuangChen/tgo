package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/G-Cool-ThanosGo/iface"
	"github.com/G-Cool-ThanosGo/models/schema"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PkgService : PKG 報表的 service
type PkgService struct{}

// PKGRequestParms : 請求PKG報表時所帶入的參數
type PKGRequestParms struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	FilePath string `json:"filePath"`
	Purpose  string `json:"purpose"`
}

// PKGData :
// 取得 Banks, CreditCardPayment, CreditCardRedeems, Gateways, GhourDecrements, GhourIncrements,
// Invoices, InvitationCode, Lots, MemberPointRedeems, Orders,
// OrderCreateRecords, OrderModRecords, ParkFeeRecords, PaymentDetails, PaymentCards, RedeemCards
type PKGData struct {
	Banks              []schema.Banks
	CreditCardPayment  []schema.CreditCardPayment
	CreditCardRedeems  []schema.CreditCardRedeems
	Gateways           []schema.Gateways
	GhourDecrements    []schema.GhourDecrements
	GhourIncrements    []schema.GhourIncrements
	Invoices           []schema.Invoices
	InvitationCode     []schema.InvitationCode
	Lots               []schema.Lots
	MemberPointRedeems []schema.MemberPointRedeems
	Orders             map[int]map[string]string
	OrderCreateRecords []schema.OrderCreateRecords
	OrderModRecords    []schema.OrderModRecords
	ParkFeeRecords     []schema.ParkingFeeRecords
	PaymentDetails     []schema.PaymentDetails
	PaymentCards       []schema.CreditCards
	RedeemCards        []schema.CreditCards
}

// HandleReportParms : 處理PKG參數
func (PkgService) HandleReportParms(c *gin.Context) (time.Time, time.Time, string, string) {
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

// GetPkgAllData : query db 取得 PKG 所需之 data
func (PkgService) GetPkgAllData(start, end time.Time, mysqldb iface.MySQLDB, rawDB *sql.DB, gdb *gorm.DB) PKGData {
	// 取得 pkg orders OrdersID AccountID LotsID ParkingRecordDetailsID
	pkgOrders, pkgExtraction := mysqldb.PkgOperator.GetByParkTime(rawDB, start, end)

	// 透過 order prop 取得 paymentDetails, invoices, invitationCode, memberPointRedeems, orderCreateRecords, orderModRecords, creditCardRedeems, paymentDetailsID, RedeemCardID
	pkgDataByOrdersID := mysqldb.PkgOperator.GetByOrdersID(pkgExtraction, gdb)

	// 透過 parking_record_details prop 取得 lots, parkFeeRecords
	lotsAndParkFee := mysqldb.PkgOperator.GetByParkingRecordDetailsID(pkgExtraction, gdb)

	// 透過 payment_details prop 取得 creditCardPayment, ghourDecrement, ghourIncID, paymentCardsID
	paymentAndUseGhour := mysqldb.PkgOperator.GetByPaymentDetailsID(pkgDataByOrdersID.PaymentDetailsID, gdb)

	// 取得加密信用卡卡號
	redeemCards := mysqldb.CreditCardsOperator.GetByCardsID(pkgDataByOrdersID.RedeemCardID, gdb)
	paymentCards := mysqldb.CreditCardsOperator.GetByCardsID(paymentAndUseGhour.PaymentCardsID, gdb)

	// 取得 收單行 銀行
	gateways := mysqldb.GatewaysOperator.GetAllGateWays(gdb)
	banks := mysqldb.BanksOperator.GetAllBank(gdb)

	pkg := PKGData{
		banks,
		paymentAndUseGhour.CreditCardPayment,
		pkgDataByOrdersID.CreditCardRedeems,
		gateways,
		paymentAndUseGhour.GhourDecrements,
		paymentAndUseGhour.GhourIncrements,
		pkgDataByOrdersID.Invoices,
		pkgDataByOrdersID.InvitationCode,
		lotsAndParkFee.Lots,
		pkgDataByOrdersID.MemberPointRedeems,
		pkgOrders,
		pkgDataByOrdersID.OrderCreateRecords,
		pkgDataByOrdersID.OrderModRecords,
		lotsAndParkFee.ParkFeeRecords,
		pkgDataByOrdersID.PaymentDetails,
		paymentCards,
		redeemCards,
	}

	return pkg
}
