package dboperator

import (
	"github.com/G-Cool-ThanosGo/models/schema"
	"github.com/jinzhu/gorm"
)

// GhourDecrementsOperator : 操作 msqdbt1.ghour_decrements
type GhourDecrementsOperator struct{}

// GhourIncrementsOperator : 操作 msqdbt1.ghour_increments
type GhourIncrementsOperator struct{}

// GetByPydID : 依照 paymentDetailsID 取得 ghourDecrements
func (GhourDecrementsOperator) GetByPydID(paymentDetailsID []int, db *gorm.DB) []schema.GhourDecrements {
	var ghourDecrements []schema.GhourDecrements
	db.Find(&ghourDecrements, "type_value in (?)", paymentDetailsID)

	return ghourDecrements
}

// GetByID : 依照 id 取得 ghourIncrements
func (GhourIncrementsOperator) GetByID(ghIncID []int, db *gorm.DB) []schema.GhourIncrements {
	var ghourIncrements []schema.GhourIncrements
	db.Find(&ghourIncrements, "id in (?)", ghIncID)

	return ghourIncrements
}
