package dboperator

import (
	"github.com/G-Cool-ThanosGo/models/schema"
	"github.com/jinzhu/gorm"
)

// OrderCreateRecordOperator : 人工成立訂單
type OrderCreateRecordOperator struct{}

// OrderModRecordOperator : 取消訂單原因
type OrderModRecordOperator struct{}

// GetByOrdersID : 用 orderID 取得人工 成立/調整 訂單
func (OrderCreateRecordOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.OrderCreateRecords {
	var orderCreateRecords []schema.OrderCreateRecords
	db.Find(&orderCreateRecords, "order_id in (?)", ordersID)

	// 若沒有 人工 成立/調整 訂單 則返回都為0的紀錄
	if len(orderCreateRecords) == 0 {
		orderCreateRecords = append(orderCreateRecords, schema.OrderCreateRecords{OrderID: 0, Reason: "null", Type: 0})
	}

	return orderCreateRecords
}

// GetByOrdersID : 用 orderID 取得 取消調整訂單
func (OrderModRecordOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.OrderModRecords {
	var orderModRecords []schema.OrderModRecords
	db.Find(&orderModRecords, "order_id in (?)", ordersID)

	// 若沒有 人工 成立/調整 訂單 則返回都為0的紀錄
	if len(orderModRecords) == 0 {
		orderModRecords = append(orderModRecords, schema.OrderModRecords{OrderID: 0, Remark: "null"})
	}

	return orderModRecords
}
