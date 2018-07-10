package dboperator

import (
	"github.com/G-Cool-ThanosGo/models/schema"
	"github.com/jinzhu/gorm"
)

// CreditCardsOperator : 操作 msqdbt1.credit_cards
type CreditCardsOperator struct{}

// CreditCardPaymentOperator : 操作 msqdbt1.credit_card_payment
type CreditCardPaymentOperator struct{}

// CreditCardRedeemsOperator : 操作 msqdbt1.credit_card_redeems
type CreditCardRedeemsOperator struct{}

// InvoicesOperator : 操作 msqdbt1.invoices
type InvoicesOperator struct{}

// PaymentDetailsOperator : 操作 msqdbt1.payment_details
type PaymentDetailsOperator struct{}

// GetByCardsID : 依照 creditCardsID 取得信用卡資訊
func (CreditCardsOperator) GetByCardsID(creditCardsID []int, db *gorm.DB) []schema.CreditCards {
	var creditCards []schema.CreditCards
	db.Find(&creditCards, "id in (?)", creditCardsID)

	return creditCards
}

// GetByPydID : 依照 paymentDetailsID 取得使用的信用卡優惠資訊
func (CreditCardPaymentOperator) GetByPydID(paymentDetailsID []int, db *gorm.DB) []schema.CreditCardPayment {
	var creditCardPayment []schema.CreditCardPayment
	db.Table("credit_card_payment").Find(&creditCardPayment, "payment_detail_id in (?)", paymentDetailsID)

	return creditCardPayment
}

// GetByOrdersID : 依照 ordersID 取得使用的信用卡優惠資訊
func (CreditCardRedeemsOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.CreditCardRedeems {
	var creditCardRedeems []schema.CreditCardRedeems
	db.Find(&creditCardRedeems, "order_id in (?)", ordersID)

	return creditCardRedeems
}

// GetByOrdersID : 依照 order_id 取得發票資訊
func (InvoicesOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.Invoices {
	var invoices []schema.Invoices
	db.Find(&invoices, "order_id in (?)", ordersID)

	return invoices
}

// GetByOrdersID : 依照 order_id 取得付款明細
func (PaymentDetailsOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.PaymentDetails {
	var paymentDetails []schema.PaymentDetails
	db.Find(&paymentDetails, "order_id in (?)", ordersID)

	return paymentDetails
}
