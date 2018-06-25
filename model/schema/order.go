package schema

import "time"

type Orders struct {
	ID            int
	Status        int
	Amount        int
	CreatedAt     time.Time
	OrderNum      string
	PaidAmount    int
	PaymentStatus int
	InvoiceStatus int
	CaptureTime   *time.Time
	RefundTime    *time.Time
	OrderNumMom   string
}
