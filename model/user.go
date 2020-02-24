package model

// User is user infomation // subはいらんぽいから外してる
type User struct {
	Common               `xorm:"extends"`
	Email                string `json:"email"`
	DisplayName          string `json:"displayName"`
	AvatarURL            string `json:"avatarUrl"`
	About                string `json:"about"`
	TotalPrice           uint64 `json:"totalPrice"`
	RecommendationBookID uint64 `json:"recommendationBookId"`
	IsTermsOfService     uint64 `json:"isTermsOfService"`
}

//UserCreateInput はユーザ作成のモデル
type UserCreateInput struct {
	Email            string `json:"email" binding:"required"`
	DisplayName      string `json:"displayName" binding:"required"`
	IsTermsOfService uint64 `json:"isTermsOfService" binding:"required"`
}

//UserUpdateInput はユーザ更新のモデル
type UserUpdateInput struct {
	DisplayName          string `json:"displayName" binding:"required"`
	AvatarURL            string `json:"avatarUrl"`
	About                string `json:"about"`
	RecommendationBookID uint64 `json:"recommendationBookId"`
	IsTermsOfService     uint64 `json:"isTermsOfService"`
}

// TableName represents db table name
func (User) TableName() string {
	return "users"
}

// DefaultImg is
const DefaultImg = "https://ay1st.s3-ap-northeast-1.amazonaws.com/images/noimage_normal.png"
