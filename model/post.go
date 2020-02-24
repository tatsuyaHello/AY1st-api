package model

// PostInput is book and action information
type PostInput struct {
	BookBody        `xorm:"extends" json:"bookData"`
	ActionInputBody []ActionInputBody `xorm:"extends" json:"action"`
}

// Post is
type Post struct {
	UserBookRegistration `xorm:"extends"`
	DisplayName          string `json:"displayName"`
	AvatarURL            string `json:"avatarUrl"`
	BookBody             `xorm:"extends" json:"bookData"`
	Action               []*Action `xorm:"extends" json:"action"`
}

// Posts is
type Posts struct {
	UserBookRegistration `xorm:"extends"`
	DisplayName          string `json:"displayName"`
	AvatarURL            string `json:"avatarUrl"`
	BookBody             `xorm:"extends"`
}

// PostOfUser はあるユーザの投稿情報
type PostOfUser struct {
	UserBookRegistration `xorm:"extends"`
	BookBody             `xorm:"extends" json:"bookData"`
	Action               []*Action `xorm:"extends" json:"action"`
}
