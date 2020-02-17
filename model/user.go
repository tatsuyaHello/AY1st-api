package model

// User is user infomation // subはいらんぽいから外してる
type User struct {
	Common               `xorm:"extends"`
	Email                string `json:"email"`
	DisplayName          string `json:"display_name"`
	AvatarURL            string `json:"avatar_url"`
	About                string `json:"about"`
	TotalPrice           uint64 `json:"total_price"`
	RecommendationBookID uint64 `json:"recommendation_book_id"`
	IsTermsOfService     uint64 `json:"is_terms_of_service"`
}

//UserCreateInput はユーザ作成のモデル
type UserCreateInput struct {
	Email       string `json:"email" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
}

//UserUpdateInput はユーザ更新のモデル
type UserUpdateInput struct {
	DisplayName          string `json:"display_name" binding:"required"`
	AvatarURL            string `json:"avatar_url"`
	About                string `json:"about"`
	RecommendationBookID uint64 `json:"recommendation_book_id"`
	IsTermsOfService     uint64 `json:"is_terms_of_service"`
}

// TableName represents db table name
func (User) TableName() string {
	return "users"
}

// DefaultImg is
const DefaultImg = "https://ay1st.s3-ap-northeast-1.amazonaws.com/images/noimage_normal.png"
