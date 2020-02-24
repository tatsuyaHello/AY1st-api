package model

// UserBookRegistration is user book registration infomation
type UserBookRegistration struct {
	Common                   `xorm:"extends"`
	UserBookRegistrationBody `xorm:"extends"`
}

// UserBookRegistrationBody is user book registration information body
type UserBookRegistrationBody struct {
	UserID            uint64 `json:"userId"`
	BookID            uint64 `json:"bookId"`
	IsActionCompleted *bool  `json:"isActionCompleted"`
}

// TableName represents db table name
func (UserBookRegistration) TableName() string {
	return "users_books_registrations"
}
