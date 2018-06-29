package dboperator

import (
	"github.com/G-Cool-ThanosGo/models/schema"
	"github.com/jinzhu/gorm"
)

// PaymentDetailsOperator : 操作 msqdbt1.payment_details
type PaymentDetailsOperator struct{}

// InvoicesOperator : 操作 msqdbt1.invoices
type InvoicesOperator struct{}

// GetByOrdersID : 依照 order_id 取得付款明細
func (PaymentDetailsOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.PaymentDetails {
	var paymentDetails []schema.PaymentDetails
	db.Find(&paymentDetails, "order_id in (?)", ordersID)

	return paymentDetails
}

// GetByOrdersID : 依照 order_id 取得發票資訊
func (InvoicesOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.Invoices {
	var invoices []schema.Invoices
	db.Find(&invoices, "order_id in (?)", ordersID)

	return invoices
}

// type creditCardOperator struct{}
// type creditCard
