package repository

import (
	"AY1st/model"
	"AY1st/util"
	"fmt"

	"github.com/go-xorm/xorm"
)

// UserIdentitiesInterface is health check (debug)
type UserIdentitiesInterface interface {
	AddIdentity(userID uint64, sub string) error
	GetUserID(sub string) (uint64, error)
	GetOne(id uint64) (*model.UserIdentity, error)
}

// UserIdentities is
type UserIdentities struct {
	engine xorm.EngineInterface
}

// NewUserIdentities initializes Users
func NewUserIdentities(engine xorm.EngineInterface) *UserIdentities {
	ui := UserIdentities{
		engine: engine,
	}
	return &ui
}

// AddIdentity はアイデンティティを追加
func (r *UserIdentities) AddIdentity(userID uint64, sub string) error {
	identity := model.UserIdentity{}
	identity.Common.UnsetDefaltCols()
	identity.UserID = userID
	identity.Sub = sub
	affected, err := r.engine.InsertOne(&identity)
	if err != nil {
		return err
	}
	if affected != 1 {
		return fmt.Errorf("failed to add identity")
	}
	return nil
}

// GetUserID sub でユーザーIDを取得
func (r *UserIdentities) GetUserID(sub string) (uint64, error) {
	userIdentity := &model.UserIdentity{}
	ok, err := r.engine.Where("sub = ?", sub).Get(userIdentity)
	if err != nil {
		util.GetLogger().Error(err)
		return 0, fmt.Errorf("can not get user identity")
	}
	if !ok {
		return 0, fmt.Errorf("can not get user identity")
	}
	return userIdentity.UserID, nil
}

// GetOne は一意なUserIdentityを取得
func (r *UserIdentities) GetOne(id uint64) (*model.UserIdentity, error) {
	userIdentity := &model.UserIdentity{}
	ok, err := r.engine.Where("user_id = ?", id).Get(userIdentity)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get user identity")
	}
	if !ok {
		return nil, fmt.Errorf("can not get user identity")
	}
	return userIdentity, nil
}
