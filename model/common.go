package model

import (
	"time"
)

// Common テーブル共通項目を定義
type Common struct {
	ID        uint64     `xorm:"pk autoincr index(pk)" json:"id"`
	CreatedAt *time.Time `xorm:"created notnull" json:"createdAt"`
	UpdatedAt *time.Time `xorm:"updated notnull" json:"updatedAt"`
}

// TableName this should not be used
func (Common) TableName() string {
	return ""
}

// UnsetDefaltCols sets init data
func (m *Common) UnsetDefaltCols() {
	m.CreatedAt = nil
	m.UpdatedAt = nil
}
