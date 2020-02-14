package model

// UserBookRegistration is user book registration infomation
type UserBookRegistration struct {
	Common                   `xorm:"extends"`
	UserBookRegistrationBody `xorm:"extends"`
}

// UserBookRegistrationBody is user book registration information body
type UserBookRegistrationBody struct {
	UserID            uint64 `json:"user_id"`
	BookID            uint64 `json:"book_id"`
	IsActionCompleted uint64 `json:"is_action_completed"`
}

// TableName represents db table name
func (UserBookRegistration) TableName() string {
	return "users_books_registrations"
}
