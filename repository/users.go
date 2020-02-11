package repository

import (
	"fmt"

	"AY1st/model"
	"AY1st/util"

	"github.com/go-pp/pp"
	"github.com/go-xorm/xorm"
)

// UsersInterface is health check (debug)
type UsersInterface interface {
	GetMe(sub string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(input *model.UserCreateInput) (*model.User, error)
	GetOne(id uint64) (*model.User, error)
	Delete(id uint64, sub string) error
	Update(id uint64, input *model.UserUpdateInput) (*model.User, error)
	GetByDisplayName(displayName string) (*model.User, error)
}

// Users is health check (debug)
type Users struct {
	engine xorm.EngineInterface
}

// NewUsers initializes Users
func NewUsers(engine xorm.EngineInterface) *Users {
	u := Users{
		engine: engine,
	}
	return &u
}

// GetMe get self
func (u *Users) GetMe(sub string) (*model.User, error) {

	session := u.engine.Table("user_identities")
	session.Select("users.*")
	session.Join("INNER", "users", "users.id = user_identities.user_id")
	session.Where("sub = ?", sub)

	me := &model.User{}
	_, err := session.Get(me)

	if err != nil {
		pp.Println(err)
		return nil, err
	}

	return me, nil
}

// GetOneByEmail はemailでユーザー情報を取得
func (u *Users) GetByEmail(email string) (*model.User, error) {
	logger := util.GetLogger()
	user := &model.User{}
	ok, err := u.engine.Where("email = ?", email).Get(user)
	if err != nil {
		logger.Error(err.Error())
		return nil, fmt.Errorf("can not get user")
	}
	if !ok {
		return nil, fmt.Errorf("can not get user")
	}
	return user, nil
}

// Create はユーザーの新規登録
func (u *Users) Create(input *model.UserCreateInput) (*model.User, error) {
	logger := util.GetLogger()

	user := &model.User{}
	user.Common.UnsetDefaltCols()
	user.Email = input.Email
	user.DisplayName = input.DisplayName
	session := u.engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	affected, err := session.InsertOne(user)
	if err != nil {
		session.Rollback()
		logger.Error(err)
		return nil, err
	}
	if affected != 1 {
		session.Rollback()
		logger.Error(err)
		return nil, fmt.Errorf("no user was created")
	}
	err = session.Commit()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return user, nil
}

// GetOne は一意なユーザーを取得
func (u *Users) GetOne(id uint64) (*model.User, error) {
	user := &model.User{}
	ok, err := u.engine.ID(id).Get(user)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get user")
	}
	if !ok {
		return nil, fmt.Errorf("can not get user")
	}
	return user, nil
}

// Delete はユーザーの削除
func (u *Users) Delete(id uint64, sub string) error {
	session := u.engine.NewSession()
	defer session.Close()
	err := session.Begin()

	userIdentity := &model.UserIdentity{}
	affected, err := u.engine.ID(id).Delete(userIdentity)
	if err != nil {
		session.Rollback()
		return err
	}
	if affected != 1 {
		session.Rollback()
		return fmt.Errorf("could not delete user-id = %v", id)
	}

	//NOTE フロントと調整
	/*
		// AWS cognito上のデータを削除
		cognitoUserName, err := infra.GetCognitoUserName(sub)
		if err != nil {
			session.Rollback()
			return err
		}

		_, err = infra.DeleteCognitoUser(cognitoUserName)
		if err != nil {
			session.Rollback()
			return err
		}
	*/

	err = session.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Update はユーザーの更新
func (u *Users) Update(id uint64, input *model.UserUpdateInput) (*model.User, error) {

	user := &model.User{}
	ok, err := u.engine.ID(id).ForUpdate().Get(user)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("no data to update")
	}

	user.DisplayName = input.DisplayName
	user.About = input.About
	user.RecommendationBook = input.RecommendationBook
	user.IsTermsOfService = input.IsTermsOfService
	now := util.GetTimeNow()
	user.Common.UpdatedAt = &now
	affected, err := u.engine.ID(id).Update(user)
	if err != nil {
		return nil, err
	}
	if affected != 1 {
		return nil, fmt.Errorf("no data affected")
	}
	return user, nil
}

// GetByDisplayName は一意なユーザーを取得
func (u *Users) GetByDisplayName(displayName string) (*model.User, error) {
	user := &model.User{}
	ok, err := u.engine.Where("display_name = ?", displayName).Get(user)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get user")
	}
	if !ok {
		return nil, fmt.Errorf("can not get user")
	}
	return user, nil
}
