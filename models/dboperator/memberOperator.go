package dboperator

import (
	"github.com/G-Cool-ThanosGo/models/schema"
	"github.com/jinzhu/gorm"
)

// InvitationCodeOperator : 操作 msqdbt1.invitationCode
type InvitationCodeOperator struct{}

// MemberPointRedeemsOperator : 操作 msqdbt1.member_point_redeems
type MemberPointRedeemsOperator struct{}

// GetByAccountID : 以 account_id 取得 邀請碼
func (InvitationCodeOperator) GetByAccountID(accountsID []int, db *gorm.DB) []schema.InvitationCode {
	var invitationCode []schema.InvitationCode

	// 若未定義 table sql 會下 from invitation_codes
	db.Table("invitation_code").Find(&invitationCode, "account_id in (?)", accountsID)

	return invitationCode
}

// GetByOrdersID : 以 order_id 取得 會員點數優惠
func (MemberPointRedeemsOperator) GetByOrdersID(ordersID []int, db *gorm.DB) []schema.MemberPointRedeems {
	var memberPointRedeems []schema.MemberPointRedeems
	db.Find(&memberPointRedeems, "order_id in (?)", ordersID)

	if len(memberPointRedeems) == 0 {
		memberPointRedeems = append(memberPointRedeems, schema.MemberPointRedeems{OrderID: 0, GatewayID: 0, Type: 0, DiscountAmount: 0, DiscountHours: 0, BonusPoint: 0})
	}

	return memberPointRedeems
}
