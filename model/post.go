package model

// PostInput is book and action information
type PostInput struct {
	BookBody        `xorm:"extends" json:"book_data"`
	ActionInputBody []ActionInputBody `xorm:"extends" json:"action"`
}

// Post is
type Post struct {
	UserBookRegistration `xorm:"extends"`
	DisplayName          string `json:"display_name"`
	AvartarURL           string `json:"avartar_url"`
	BookBody             `xorm:"extends" json:"book_data"`
	Action               []*Action `xorm:"extends" json:"action"`
}

// Posts is
type Posts struct {
	UserBookRegistration `xorm:"extends"`
	DisplayName          string `json:"display_name"`
	AvartarURL           string `json:"avartar_url"`
	BookBody             `xorm:"extends"`
}

// PostOfUser はあるユーザの投稿情報
type PostOfUser struct {
	UserBookRegistration `xorm:"extends"`
	BookBody             `xorm:"extends" json:"book_data"`
	Action               []*Action `xorm:"extends" json:"action"`
}
