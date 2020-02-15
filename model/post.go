package model

// PostInput is book and action information
type PostInput struct {
	BookBody        `xorm:"extends"`
	ActionInputBody []ActionInputBody `xorm:"extends" json:"action_input_body"`
}

// Post is
type Post struct {
	UserBookRegistration `xorm:"extends"`
	DisplayName          string `json:"display_name"`
	AvartarURL           string `json:"avartar_url"`
	BookBody             `xorm:"extends"`
	ActionBody           []*ActionBody `xorm:"extends" json:"action_body"`
}
