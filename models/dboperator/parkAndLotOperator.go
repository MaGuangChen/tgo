package dboperator

import (
	"github.com/G-Cool-ThanosGo/models/schema"
	"github.com/jinzhu/gorm"
)

// ParkingRecordDetailsOperator :
type ParkingRecordDetailsOperator struct{}

// LotsOperator : 操作 msqdbt1.lots 表單
type LotsOperator struct{}

// ParkingFeeRecordsOperator : 操作 msqdbt1.parking_fee_records 表單
type ParkingFeeRecordsOperator struct{}

// GetAllByID : 透過 lot.id 取得車廠資訊
func (LotsOperator) GetAllByID(lotID []int, db *gorm.DB) []schema.Lots {
	var lots []schema.Lots
	db.Find(&lots, "id in (?)", lotID)

	return lots
}

// GetAllByPrdID : 透過 parkingRecordDetailsID 取得 停車計費明細
func (ParkingFeeRecordsOperator) GetAllByPrdID(parkingRecordDetailsID []int, db *gorm.DB) []schema.ParkingFeeRecords {
	var parkingFeeRecords []schema.ParkingFeeRecords
	db.Find(&parkingFeeRecords, "id in (?)", parkingRecordDetailsID)

	return parkingFeeRecords
}
