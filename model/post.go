package model

// PostInput is book and action information
type PostInput struct {
	BookBody        `xorm:"extends"`
	ActionInputBody []ActionInputBody `xorm:"extends" json:"action_input_body"`
}
