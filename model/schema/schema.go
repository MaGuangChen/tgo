package schema

import "time"

type Orders struct {
	ID                   int `gorm:"primary_key"`
	Status               int
	Amount               int
	CreatedAt            time.Time
	OrderNum             string
	PaidAmount           int
	PaymentStatus        int
	InvoiceStatus        int
	CaptureTime          *time.Time
	RefundTime           *time.Time
	OrderNumMom          string
	AccountID            int
	ParkingRecordDetails ParkingRecordDetails
}

type ParkingRecordDetails struct {
	ID              int `gorm:"primary_key"`
	OrderID         int
	ParkingRecordID int
}

type ParkingRecord struct {
	ID         int `gorm:"primary_key"`
	EnteredAt  *time.Time
	ExitedAt   *time.Time
	LotCode    string
	EntryType  int
	EntryValue string
}

type PaymentDetails struct {
}

type invoices struct {
}
