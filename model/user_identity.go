package model

// UserIdentity is user identity infomation
type UserIdentity struct {
	Common `xorm:"extends"`
	Sub    string `json:"sub"`
	UserID uint64 `json:"userId"`
}

// TableName represents db table name
func (UserIdentity) TableName() string {
	return "user_identities"
}
