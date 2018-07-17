package dboperator

import (
	"sort"

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

// GetByCardsID : 依照 creditCardsID 取得信用卡資訊 sorted by ID
func (CreditCardsOperator) GetByCardsID(creditCardsID []int, db *gorm.DB) []schema.CreditCards {
	var creditCards []schema.CreditCards
	db.Find(&creditCards, "id in (?)", creditCardsID)
	sort.SliceStable(creditCards, func(i, j int) bool {
		return creditCards[i].ID < creditCards[j].ID
	})

	return creditCards
}

// GetByPydID : 依照 paymentDetailsID 取得使用的信用卡優惠資訊 sorted by PaymentDetailID
func (CreditCardPaymentOperator) GetByPydID(paymentDetailsID []int, db *gorm.DB) []schema.CreditCardPayment {
	var creditCardPayment []schema.CreditCardPayment
	db.Table("credit_card_payment").Find(&creditCardPayment, "payment_detail_id in (?)", paymentDetailsID)
	sort.SliceStable(creditCardPayment, func(i, j int) bool {
		return creditCardPayment[i].PaymentDetailID < creditCardPayment[j].PaymentDetailID
	})

	return creditCardPayment
}

// GetByOrdersID : 依照 ordersID 取得使用的信用卡優惠資訊
func (CreditCardRedeemsOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.CreditCardRedeems {
	var creditCardRedeems []schema.CreditCardRedeems
	db.Find(&creditCardRedeems, "order_id in (?)", ordersID)
	sort.SliceStable(creditCardRedeems, func(i, j int) bool {
		return creditCardRedeems[i].PaymentDetailID < creditCardRedeems[j].PaymentDetailID
	})

	return creditCardRedeems
}

// GetByOrdersID : 依照 order_id 取得發票資訊 sorted by OrderID
func (InvoicesOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.Invoices {
	var invoices []schema.Invoices
	db.Find(&invoices, "order_id in (?)", ordersID)
	sort.SliceStable(invoices, func(i, j int) bool {
		return invoices[i].OrderID < invoices[j].OrderID
	})

	return invoices
}

// GetByOrdersID : 依照 order_id 取得付款明細 sorted by OrderID
func (PaymentDetailsOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.PaymentDetails {
	var paymentDetails []schema.PaymentDetails
	db.Find(&paymentDetails, "order_id in (?)", ordersID)
	sort.SliceStable(paymentDetails, func(i, j int) bool {
		return paymentDetails[i].OrderID < paymentDetails[j].OrderID
	})

	return paymentDetails
}
