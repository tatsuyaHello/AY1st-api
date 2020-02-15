package model

// User is user infomation // subはいらんぽいから外してる
type User struct {
	Common             `xorm:"extends"`
	Email              string `json:"email"`
	DisplayName        string `json:"display_name"`
	AvartarURL         string `json:"avartar_url"`
	About              string `json:"about"`
	TotalPrice         uint64 `json:"total_price"`
	RecommendationBook string `json:"recommendation_book"`
	IsTermsOfService   uint64 `json:"is_terms_of_service"`
}

//UserCreateInput はユーザ作成のモデル
type UserCreateInput struct {
	Email       string `json:"email" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
}

//UserUpdateInput はユーザ更新のモデル
type UserUpdateInput struct {
	DisplayName        string `json:"display_name" binding:"required"`
	About              string `json:"about"`
	RecommendationBook string `json:"recommendation_book"`
	IsTermsOfService   uint64 `json:"is_terms_of_service"`
}

// TableName represents db table name
func (User) TableName() string {
	return "users"
}
