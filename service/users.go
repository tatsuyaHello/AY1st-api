package service

import (
	"AY1st/model"
	"AY1st/repository"
	"fmt"
)

// UsersInterface is
type UsersInterface interface {
	GetByEmail(email string) (*model.User, error)
	GetUserID(sub string) (uint64, error)
	AddIdentity(userID uint64, sub string) error
	GetMe(sub string) (*model.User, error)
	Create(user *model.UserCreateInput) (*model.User, error) //TODO 戻り値を適切な構造体にする
	GetOne(id uint64) (*model.User, error)
	Delete(id uint64) error
	Update(id uint64, user *model.UserUpdateInput) (*model.User, error)
}

// Users is
type Users struct {
	UsersRepo          repository.UsersInterface
	UserIdentitiesRepo repository.UserIdentitiesInterface
}

// NewUsers is
func NewUsers(UsersRepo repository.UsersInterface, UserIdentitiesRepo repository.UserIdentitiesInterface) *Users {
	u := Users{
		UsersRepo:          UsersRepo,
		UserIdentitiesRepo: UserIdentitiesRepo,
	}
	return &u
}

// GetByEmail はユーザーを一意識別名を元に取得
func (u *Users) GetByEmail(email string) (*model.User, error) {
	return u.UsersRepo.GetByEmail(email)
}

// GetUserID はユーザーアイデンティティに紐づくUserIDを取得
func (u *Users) GetUserID(sub string) (uint64, error) {
	return u.UserIdentitiesRepo.GetUserID(sub)
}

// AddIdentity はユーザーアイデンティティを追加します
func (u *Users) AddIdentity(userID uint64, sub string) error {
	return u.UserIdentitiesRepo.AddIdentity(userID, sub)
}

// GetMe is
func (u *Users) GetMe(sub string) (*model.User, error) {
	return u.UsersRepo.GetMe(sub)
}

// Create はユーザー新規登録
func (u *Users) Create(input *model.UserCreateInput) (*model.User, error) {
	us, err := u.UsersRepo.GetByEmail(input.Email)
	if us != nil || err == nil {
		return nil, model.NewError(model.ErrorUserSubDuplicate, "user already exists")
	}

	return u.UsersRepo.Create(input)
}

// GetOne は一意なユーザーを取得
func (u *Users) GetOne(id uint64) (*model.User, error) {
	user, err := u.UsersRepo.GetOne(id)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "user not found")
	}

	return user, nil
}

// Update はユーザーを更新
func (u *Users) Update(id uint64, input *model.UserUpdateInput) (*model.User, error) {
	_, found := u.UsersRepo.GetOne(id)
	if found != nil {
		return nil, fmt.Errorf("user does not exist")
	}

	return u.UsersRepo.Update(id, input)
}

// Delete はユーザー削除
func (u *Users) Delete(id uint64) error {
	_, err := u.UsersRepo.GetOne(id)
	if err != nil {
		return model.NewError(model.ErrorResourceNotFound, "user not found")
	}

	userIdentity, err := u.UserIdentitiesRepo.GetOne(id)
	if err != nil {
		return model.NewError(model.ErrorResourceNotFound, "user identity not found")
	}

	// ユーザーの削除
	err = u.UsersRepo.Delete(id, userIdentity.Sub)
	if err != nil {
		return err
	}

	return nil
}
