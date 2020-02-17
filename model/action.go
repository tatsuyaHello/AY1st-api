package model

// Action is action infomation
type Action struct {
	Common     `xorm:"extends"`
	ActionBody `xorm:"extends"`
}

// ActionBody is
type ActionBody struct {
	IsFinished             *bool  `json:"is_finished"`
	UserBookRegistrationID uint64 `json:"user_book_registration_id"`
	Content                string `json:"content"`
}

// ActionInputBody is
type ActionInputBody struct {
	Content string `json:"content"`
}

//ActionUpdateInput はユーザ更新のモデル
type ActionUpdateInput struct {
	ID         uint64 `xorm:"pk autoincr index(pk)" json:"id"`
	IsFinished *bool  `json:"is_finished"`
	Content    string `json:"content"`
}

// TableName represents db table name
func (Action) TableName() string {
	return "actions"
}

// TableName represents db table name
func (ActionBody) TableName() string {
	return "actions"
}
