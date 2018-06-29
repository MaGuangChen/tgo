package schema

import "time"

// PaymentDetails : 訂單付款明細
type PaymentDetails struct {
	ID      int `gorm:"primary_key"`
	OrderID int
	Type    int
	Amount  float64
	More    string
}

// Invoices : 發票
type Invoices struct {
	OrderID   int
	Num       string
	CreatedAt *time.Time
	Vat       string
	Status    int
}

// InvitationCode : 邀請碼
type InvitationCode struct {
	AccountID int
	Code      string
}

// MemberPointRedeems : 會員點數優惠
type MemberPointRedeems struct {
	OrderID        int
	GatewayID      int
	Type           int
	DiscountAmount float64
	DiscountHours  float64
	BonusPoint     float64
}
