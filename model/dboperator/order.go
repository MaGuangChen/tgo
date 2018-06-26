package dboperator

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/G-Cool-ThanosGo/model/schema"
	"github.com/G-Cool-ThanosGo/util"
	"github.com/jinzhu/gorm"
)

type op interface {
	byTime()
}

type OrderOperator struct{}
type ParkingRecordOperator struct{}
type ParkingRecordDetailOperator struct{}

func (ParkingRecordOperator) GetByExitedAt(d *gorm.DB, start time.Time, end time.Time) []schema.ParkingRecord {
	defer d.Close()
	var parkingRecord []schema.ParkingRecord
	d.Find(&parkingRecord, "exited_at >= ? AND exited_at <= ?", start, end)

	return parkingRecord
}

func (ParkingRecordDetailOperator) GetByParkingRecordID(d *gorm.DB, prID []int) []schema.ParkingRecordDetails {
	defer d.Close()
	var parkingRecordDetails []schema.ParkingRecordDetails
	d.Find(&parkingRecordDetails, "parking_record_id in (?)", prID)

	return parkingRecordDetails
}

func (OrderOperator) GetByID(d *gorm.DB, id []int) []schema.Orders {
	defer d.Close()
	var orders []schema.Orders
	d.Find(&orders, "id = ?", id)

	return orders
}

type t map[string]string

func (OrderOperator) GetByParkTime(rawDB *sql.DB, s time.Time, e time.Time) []t {
	sqlSynx := `SELECT o.* from orders o
	INNER JOIN msqdbt1.parking_record_details prd on prd.order_id = o.id
	INNER JOIN msqdbt1.parking_records pr on pr.id = prd.parking_record_id
	WHERE pr.exited_at >= ` + "'" + s.UTC().Format("2006-01-02 15:04:05") + "'" + " AND pr.exited_at <= " + "'" + e.UTC().Format("2006-01-02 15:04:05") + "'"
	result := ScanAndGetResult(rawDB, sqlSynx)

	return result
}

func ScanAndGetResult(rawDB *sql.DB, sqlSynx string) []t {
	rows, queryErr := rawDB.Query(sqlSynx)
	util.CheckError(queryErr)

	cols, _ := rows.Columns()
	values := make([]sql.RawBytes, len(cols))
	scans := make([]interface{}, len(cols))
	results := make([]t, 2000, 2000)

	for i := range values {
		scans[i] = &values[i]
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		if queryErr := rows.Scan(scans...); queryErr != nil {
			fmt.Println("Error")
		}
		row := make(map[string]string)
		for j, v := range values {
			key := cols[j]
			row[key] = string(v)
		}
		results[i] = row
		i++
	}

	return results
}
