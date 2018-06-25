package dboperator

import (
	"time"

	"github.com/G-Cool-ThanosGo/model/schema"
	"github.com/jinzhu/gorm"
)

type op interface {
	byTime()
}

type OrderOperator struct{}
type ParkingRecordOperator struct{}
type ParkingRecordDetailOperator struct{}

func (ParkingRecordOperator) GetByExitedAt(d *gorm.DB, start time.Time, end time.Time) []schema.ParkingRecord {
	var parkingRecord []schema.ParkingRecord
	d.Find(&parkingRecord, "exited_at >= ? AND exited_at <= ?", start, end)

	return parkingRecord
}

func (ParkingRecordDetailOperator) GetByParkingRecordID(d *gorm.DB, prID []int) []schema.ParkingRecordDetails {
	var parkingRecordDetails []schema.ParkingRecordDetails
	d.Find(&parkingRecordDetails, "parking_record_id in (?)", prID)

	return parkingRecordDetails
}

func (OrderOperator) GetByID(d *gorm.DB, id int) []schema.Orders {
	var orders []schema.Orders
	d.Find(&orders, "id = ?", id)

	return orders
}
