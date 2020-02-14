package model

// Book is book infomation
type Book struct {
	Common   `xorm:"extends"`
	BookBody `xorm:"extends"`
}

// BookBody is
type BookBody struct {
	RakutenID  string `json:"rakuten_id"`
	Title      string `json:"title"`
	Price      uint64 `json:"price"`
	Author     string `json:"author"`
	BookImgURL string `json:"book_img_url"`
}

// BookActionInput is book and action information
type BookActionInput struct {
	BookBody        `xorm:"extends"`
	ActionInputBody []ActionInputBody `xorm:"extends" json:"action_input_body"`
}

// TableName represents db table name
func (Book) TableName() string {
	return "books"
}
