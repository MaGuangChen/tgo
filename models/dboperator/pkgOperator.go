package dboperator

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// PkgOperator : 操作 msqdbt1.orders
type PkgOperator struct{}

// RecordType : 下 raw sql 時返回的 record type
type RecordType map[int]map[string]string

// PKGOrdersExtractProp : 將pkg
type PKGOrdersExtractProp struct {
	OrdersID               []int
	AccountID              []int
	LotsID                 []int
	ParkingRecordDetailsID []int
}

// PKGCommonQuery : 出PKG報表的 common query
func (PkgOperator) pkgCommonQuery() string {
	return `SELECT 
	o.status, o.amount, o.created_at, o.order_num, o.paid_amount, o.payment_status, o.invoice_status, o.account_id,
	prd.id, prd.order_id, prd.lot_id, pr.magic_no
	FROM orders o 
	JOIN msqdbt1.parking_record_details prd ON prd.order_id = o.id
	JOIN msqdbt1.parking_records pr ON pr.id = prd.parking_record_id
	`
}

// 從 orders 中 將要進行下次 query 的 properties 取出
func (PkgOperator) pkgOrdersExtract(pkgAndOrders RecordType) PKGOrdersExtractProp {
	ordersID := make([]int, len(pkgAndOrders))
	accountsID := make([]int, len(pkgAndOrders))
	lotsID := make([]int, len(pkgAndOrders))
	parkingRecordDetailsID := make([]int, len(pkgAndOrders))

	for i, value := range pkgAndOrders {
		orderID, _ := strconv.Atoi(value["order_id"])
		ordersID[i] = orderID

		accountID, _ := strconv.Atoi(value["account_id"])
		accountsID[i] = accountID

		lotID, _ := strconv.Atoi(value["lot_id"])
		lotsID[i] = lotID

		prdID, _ := strconv.Atoi(value["id"])
		parkingRecordDetailsID[i] = prdID
	}

	pkgExtraction := PKGOrdersExtractProp{ordersID, accountsID, lotsID, parkingRecordDetailsID}

	return pkgExtraction
}

// GetByParkTime : 依照出場時間尋找訂單, 符合 PKG 規則
func (pkg *PkgOperator) GetByParkTime(rawDB *sql.DB, s time.Time, e time.Time) (RecordType, PKGOrdersExtractProp) {
	// 組 sql
	pkgQuery := pkg.pkgCommonQuery()
	condition := ` WHERE o.status IN (1, 2, 3, 4)
	AND pr.exited_at >= "` + s.UTC().Format("2006-01-02 15:04:05") + `" AND pr.exited_at <= "` + e.UTC().Format("2006-01-02 15:04:05") + `"`
	sqlSynx := pkgQuery + condition

	pkgAndOrders := ScanAndGetResult(rawDB, sqlSynx)
	pkgExtraction := pkg.pkgOrdersExtract(pkgAndOrders)

	return pkgAndOrders, pkgExtraction
}

// GetT4ManulAndRefund : 尋找建立時間在本月，且出場時間在上月的訂單
func (pkg *PkgOperator) GetT4ManulAndRefund(rawDB *sql.DB, s time.Time, e time.Time, ordersID map[int]string) RecordType {
	t4End := e.AddDate(0, 0, 3)
	prevMonthStart := s.AddDate(0, -1, 0)
	pkgQuery := pkg.pkgCommonQuery()
	condition1 := ` WHERE o.status IN (1, 2, 3, 4) AND o.created_at >= "` + s.UTC().Format("2006-01-02 15:04:05") + `" AND o.created_at <= "` + t4End.UTC().Format("2006-01-02 15:04:05") + `"`
	condition2 := ` AND pr.exited_at >= "` + prevMonthStart.UTC().Format("2006-01-02 15:04:05") + `" AND pr.exited_at <= "` + e.UTC().Format("2006-01-02 15:04:05") + `"`
	ordersIDString := ``
	for i, value := range ordersID {
		if i == 0 {
			ordersIDString = value
		}
		ordersIDString = ordersIDString + ", " + value
	}
	condition3 := ` AND o.id NOT IN (` + ordersIDString + ")"

	sqlSynx := pkgQuery + condition1 + condition3 + condition2

	result := ScanAndGetResult(rawDB, sqlSynx)

	return result
}

// GetPkgFinance : 取得PKG報表所需要的財物相關資料
func (PkgOperator) GetPkgFinance(pkgExtraction PKGOrdersExtractProp, gdb *gorm.DB) {
	pydOperator := PaymentDetailsOperator{}
	paymentDetails := pydOperator.GetByOrdersID(pkgExtraction.OrdersID, gdb)

	invoiceOperator := InvoicesOperator{}
	invoices := invoiceOperator.GetByOrdersID(pkgExtraction.OrdersID, gdb)

	invitationCodeOperator := InvitationCodeOperator{}
	invitationCode := invitationCodeOperator.GetByAccountID(pkgExtraction.AccountID, gdb)

	memberPointRedeemsOperator := MemberPointRedeemsOperator{}
	memberPointRedeems := memberPointRedeemsOperator.GetByOrdersID(pkgExtraction.OrdersID, gdb)

	fmt.Println(paymentDetails[0])
	fmt.Println(invoices[0])
	fmt.Println(invitationCode[0])
	fmt.Println("memberPointRedeems in pkg op: ", memberPointRedeems)
}
