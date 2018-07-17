package dboperator

import (
	"github.com/G-Cool-ThanosGo/models/schema"
	"github.com/jinzhu/gorm"
)

// BanksOperator :
type BanksOperator struct{}

// GatewaysOperator :
type GatewaysOperator struct{}

// GetAllBank : 取得 所有銀行
func (BanksOperator) GetAllBank(db *gorm.DB) []schema.Banks {
	var banks []schema.Banks
	db.Find(&banks)

	return banks
}

// GetAllGateWays : 取得 所有收單行
func (GatewaysOperator) GetAllGateWays(db *gorm.DB) []schema.Gateways {
	var gateways []schema.Gateways
	db.Find(&gateways)

	return gateways
}
