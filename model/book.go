package model

// Book is book infomation
type Book struct {
	Common   `xorm:"extends"`
	BookBody `xorm:"extends"`
}

// BookBody is
type BookBody struct {
	Title         string `json:"title"`
	Price         uint64 `json:"price"`
	Author        string `json:"author"`
	BookImgURL    string `json:"bookImgUrl"`
	RakutenURL    string `json:"rakutenUrl"`
	RakutenReview uint64 `json:"rakutenReview"`
	Isbn          uint64 `json:"isbn"`
}

// TableName represents db table name
func (Book) TableName() string {
	return "books"
}
