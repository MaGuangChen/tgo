package schema

import "time"

// CreditCards : 信用卡資訊
type CreditCards struct {
	ID               int `gorm:"primary_key"`
	BankCode         int
	CardNumberPublic string
}

// CreditCardRedeems : 信用卡優惠資訊
type CreditCardRedeems struct {
	ID              int `gorm:"primary_key"`
	PaymentDetailID int
	AuthCode        string
	CreditCardID    int
	BonusPoint      float64
	RecordType      int
	OrderNum        string
}

// CreditCardPayment : 信用卡付款資訊
type CreditCardPayment struct {
	PaymentDetailID int
	GatewayID       int
	CreditCardID    int
	AuthCode        string
	OrderNum        string
}

// GhourDecrements : 酷時減少
type GhourDecrements struct {
	GhourIncrementID int
	Reason           string
	TypeValue        int
	Amount           float64
}

// GhourIncrements : 酷時增加
type GhourIncrements struct {
	ID        int `gorm:"primary_key"`
	CreatedAt *time.Time
	AccountID int
	Amount    float64
	StartDate *time.Time
	Reason    string
}

// Lots : 車廠資訊
type Lots struct {
	ID           int `gorm:"primary_key"`
	Code         string
	Name         string
	VendorCode   string
	IsServiceAPI int
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

// OrderCreateRecords : 人工 成立/調整 訂單
type OrderCreateRecords struct {
	OrderID int
	Reason  string
	Type    int
}

// OrderModRecords : 取消調整訂單
type OrderModRecords struct {
	OrderID int
	Remark  string
}

// PaymentDetails : 訂單付款明細
type PaymentDetails struct {
	ID      int `gorm:"primary_key"`
	OrderID int
	Type    int
	Amount  float64
	More    string
}

// ParkingFeeRecords : 停車費用紀錄 (停車時數、實際付費時數、最高停車上限)
type ParkingFeeRecords struct {
	ParkingRecordDetailID int
	PaidHours             float64
	Hours                 float64
	Collection            string
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
